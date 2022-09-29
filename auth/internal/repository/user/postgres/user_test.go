package postgres

import (
	"auth/internal/domain/user"
	"auth/internal/domain/user/email"
	"auth/internal/domain/user/name"
	"auth/internal/domain/user/pass"
	"auth/internal/domain/user/phone"
	"auth/internal/domain/user/surname"
	"auth/pkg/docker"
	"auth/pkg/helpers"
	"auth/pkg/psql"
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var (
	repository          *Repository
	postgresContainerID string

	salt = "WcrMaOZhpYrpk79vFk"

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

func TestCreateUser(t *testing.T) {
	req := require.New(t)
	user := createTestUser()

	t.Run("create user success", func(t *testing.T) {
		err := repository.CreateUser(context.Background(), user)
		req.Empty(err)
	})

	t.Run("duplicate key error", func(t *testing.T) {
		err := repository.CreateUser(context.Background(), user)
		req.ErrorContains(err, ErrDuplicateKey.Error())
	})

}

func TestUpdateUser(t *testing.T) {
	req := require.New(t)
	userForUpdate := createTestUser()

	err := repository.CreateUser(context.Background(), userForUpdate)
	if err != nil {
		return
	}

	pass, _ := pass.NewPass("Qwerty123!", salt)

	newUser, _ := user.New(name.Name("updname"), surname.Surname("updsurname"), phone.Phone("79111111111"), *pass, email.Email("upemail@mail.ru"), user.Root)

	updateFunc := func(oldUser *user.User) (*user.User, error) {
		return user.NewWithID(oldUser.ID(), oldUser.CreatedAt(), time.Now().UTC(), newUser.Name(), newUser.Surname(), newUser.Phone(), newUser.Pass(), newUser.Email(), newUser.Role())
	}

	t.Run("update company success", func(t *testing.T) {

		res, err := repository.UpdateUser(context.Background(), userForUpdate.ID(), updateFunc)
		req.Empty(err)
		req.Equal(res.Name(), newUser.Name())
		req.Equal(res.Surname(), newUser.Surname())
		req.Equal(res.Phone(), newUser.Phone())
		req.Equal(res.Email(), newUser.Email())
		req.Equal(res.Role(), newUser.Role())
	})

	t.Run("not found", func(t *testing.T) {
		_, err := repository.UpdateUser(context.Background(), uuid.New(), updateFunc)
		req.ErrorContains(err, ErrEmptyResult.Error())
	})
}

func TestDeleteUser(t *testing.T) {
	req := require.New(t)
	user := createTestUser()
	err := repository.CreateUser(context.Background(), user)
	if err != nil {
		log.Fatalf("error creating test user: %s", err.Error())
	}

	t.Run("delete user success", func(t *testing.T) {
		err := repository.DeleteUser(context.Background(), uuid.UUID(user.ID()))
		req.Empty(err)
	})

	t.Run("not found", func(t *testing.T) {
		err := repository.DeleteUser(context.Background(), uuid.UUID(user.ID()))
		req.ErrorContains(err, ErrNotFound.Error())
	})

}

func TestReadUserByID(t *testing.T) {
	req := require.New(t)
	user := createTestUser()
	err := repository.CreateUser(context.Background(), user)
	if err != nil {
		log.Fatalf("error creating test user: %s", err.Error())
	}

	t.Run("success", func(t *testing.T) {
		userR, err := repository.ReadUserByID(context.Background(), uuid.UUID(user.ID()))
		req.Empty(err)
		req.Equal(user.ID(), userR.ID())
	})

	t.Run("not found", func(t *testing.T) {
		_, err := repository.ReadUserByID(context.Background(), uuid.New())
		req.ErrorContains(err, ErrEmptyResult.Error())
	})

}

func createTestUser() *user.User {
	name, _ := name.NewName("user")
	surname, _ := surname.NewSurname("surname")
	phone, _ := phone.NewPhone("+7 (900) 000-00-00")
	email, _ := email.NewEmail("test01@mail.ru")
	pass, _ := pass.NewPass("Password123!", salt)
	role := user.Administrator
	testUser, _ := user.New(*name, *surname, *phone, *pass, *email, role)
	return testUser
}
