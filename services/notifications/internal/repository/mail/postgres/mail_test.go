package postgres

import (
	"context"
	"fmt"
	"log"
	"notifications/internal/domain"
	"notifications/internal/domain/mail"
	"notifications/internal/domain/mail/email"
	"notifications/internal/domain/mail/message"
	"notifications/internal/domain/mail/subject"
	"notifications/pkg/docker"
	"notifications/pkg/helpers"
	"notifications/pkg/psql"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var (
	repository          *Repository
	postgresContainerID string

	postgresContainer = docker.ContainerOptions{
		Image:       "postgres",
		Name:        "postgres_testing",
		Host:        "localhost",
		InsidePort:  5432,
		ExposedPort: 5432,
		Envs: map[string]string{
			"POSTGRES_USER":     "admin",
			"POSTGRES_PASSWORD": "pass",
			"POSTGRES_DB":       "test",
		},
	}

	DSN = fmt.Sprintf("postgres://%s:%s@localhost:%d/%s?sslmode=disable", postgresContainer.Envs["POSTGRES_USER"], postgresContainer.Envs["POSTGRES_PASSWORD"], postgresContainer.ExposedPort, postgresContainer.Envs["POSTGRES_DB"])
)

func TestMain(m *testing.M) {
	//create testing postgres container
	dockerCli, err := docker.New()
	if err != nil {
		log.Fatalf("cannot create dockerCli: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*30))
	defer cancel()

	dockerCli.DownloadImage(ctx, postgresContainer.Image)
	dockerCli.DeleteContainer(ctx, postgresContainer.Name)

	postgresContainerID, err = dockerCli.CreateContainer(ctx, postgresContainer)
	if err != nil {
		log.Fatalf("cannot create testing postgres container: %s", err.Error())
	}

	running, err := dockerCli.IsContainerRunning(ctx, postgresContainerID)
	if err != nil || !running {
		log.Fatal("postgres container is not started")
	}

	//to wait for container initialization
	time.Sleep(time.Duration(time.Second))

	//connect to postgres
	pg, err := psql.New(DSN, psql.ConnAttempts(5))
	if err != nil {
		log.Fatalf("cannot connect to db: %s", err.Error())
	}

	err = helpers.MigrationsUP(DSN, "file://../../../../migrations")
	if err != nil {
		log.Fatalf("migrations up failure: %s", err.Error())
	}

	repository, err = New(pg, Options{})
	if err != nil {
		log.Fatalf("error initializing repository: %s", err.Error())
	}

	os.Exit(m.Run())
}

func TestCreateMail(t *testing.T) {
	req := require.New(t)

	mail := createTestMail()

	t.Run("create mail success", func(t *testing.T) {
		err := repository.CreateMail(context.Background(), mail)
		req.Empty(err)
	})
}

func TestUpdateMail(t *testing.T) {
	req := require.New(t)

	mailForUpdate := createTestMail()

	err := repository.CreateMail(context.Background(), mailForUpdate)
	if err != nil {
		return
	}

	email, _ := email.NewEmail("test_upd@mail.ru")
	subject, _ := subject.NewSubject("updated subject")
	message, _ := message.NewMessage("updated mesage")

	newMessage, _ := mail.NewWithID(mailForUpdate.ID(), mailForUpdate.CreatedAt(), time.Now(), *email, *subject, *message, domain.Sended)

	updateFunc := func(oldMessage *mail.Mail) (*mail.Mail, error) {
		return mail.NewWithID(oldMessage.ID(), oldMessage.CreatedAt(), newMessage.ModifiedAt(), newMessage.Recipient(), newMessage.Subject(), newMessage.Message(), newMessage.Status())
	}

	t.Run("mail update succees", func(t *testing.T) {
		res, err := repository.UpdateMail(context.Background(), mailForUpdate.ID(), updateFunc)
		req.Empty(err)
		req.Equal(res.ID(), newMessage.ID())
		req.Equal(res.ModifiedAt(), newMessage.ModifiedAt())
		req.Equal(res.Subject(), newMessage.Subject())
		req.Equal(res.Message(), newMessage.Message())
		req.Equal(res.Recipient(), newMessage.Recipient())
		req.Equal(res.Status(), newMessage.Status())
	})

	t.Run("not found", func(t *testing.T) {
		res, err := repository.UpdateMail(context.Background(), uuid.New(), updateFunc)
		req.Empty(res)
		req.ErrorContains(err, ErrEmptyResult.Error())
	})
}

func TestDeleteMail(t *testing.T) {
	req := require.New(t)

	testMail := createTestMail()

	err := repository.CreateMail(context.Background(), testMail)
	if err != nil {
		return
	}

	t.Run("success delete mail", func(t *testing.T) {
		err := repository.DeleteMail(context.Background(), testMail.ID())
		req.Empty(err)
	})

	t.Run("error delete mail", func(t *testing.T) {
		repository.DeleteMail(context.Background(), testMail.ID())
	})
}

func TestReadMailByID(t *testing.T) {
	req := require.New(t)

	testMail := createTestMail()

	err := repository.CreateMail(context.Background(), testMail)
	if err != nil {
		return
	}

	t.Run("success read mail", func(t *testing.T) {
		resp, err := repository.ReadMailByID(context.Background(), testMail.ID())
		req.Empty(err)
		req.NotEmpty(resp)
	})

	t.Run("not found", func(t *testing.T) {
		resp, err := repository.ReadMailByID(context.Background(), uuid.New())
		req.ErrorContains(err, ErrEmptyResult.Error())
		req.Empty(resp)
	})
}

func TestReadMailFiltredList(t *testing.T) {
	req := require.New(t)

	testMail := createTestMail()

	err := repository.CreateMail(context.Background(), testMail)
	if err != nil {
		return
	}

	t.Run("success read mail's list", func(t *testing.T) {
		resp, err := repository.ReadMailFiltredList(context.Background(), map[string]interface{}{"status": domain.NotSended, "subject": "test%"})
		req.Empty(err)
		req.NotEmpty(resp)
	})
}

func createTestMail() *mail.Mail {
	recipient, _ := email.NewEmail("kanya384@mail.ru")
	subject, _ := subject.NewSubject("test subject")
	message, _ := message.NewMessage("test message")

	mail, _ := mail.New(*recipient, *subject, *message)
	return mail
}
