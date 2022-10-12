package postgres

import (
	"companies/internal/domain/company"
	"companies/internal/domain/company/inn"
	"companies/internal/domain/company/name"
	"companies/pkg/docker"
	"companies/pkg/helpers"
	"companies/pkg/psql"
	"companies/pkg/types/address"
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

var companiesList []*company.Company

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

	//initialize test company

	i := 6
	for i > 0 {
		companiesList = append(companiesList, createTestCompany())
		i--
	}

	os.Exit(m.Run())
}

func TestCreateCompany(t *testing.T) {
	req := require.New(t)

	t.Run("create company success", func(t *testing.T) {
		err := repository.CreateCompany(context.Background(), companiesList[0])
		req.Empty(err)
	})

	t.Run("duplicate key error", func(t *testing.T) {
		err := repository.CreateCompany(context.Background(), companiesList[0])
		req.ErrorContains(err, ErrDuplicateKey.Error())
	})

}

func TestUpdateCompany(t *testing.T) {
	req := require.New(t)
	err := repository.CreateCompany(context.Background(), companiesList[1])
	if err != nil {
		return
	}

	t.Run("update company success", func(t *testing.T) {
		err := repository.UpdateCompany(context.Background(), companiesList[1])
		req.Empty(err)
	})

	t.Run("not found", func(t *testing.T) {
		err := repository.UpdateCompany(context.Background(), companiesList[2])
		req.ErrorContains(err, ErrNotFound.Error())
	})

}

func TestDeleteCompany(t *testing.T) {
	req := require.New(t)
	err := repository.CreateCompany(context.Background(), companiesList[3])
	if err != nil {
		return
	}

	t.Run("delete company success", func(t *testing.T) {
		err := repository.DeleteCompany(context.Background(), companiesList[3].ID())
		req.Empty(err)
	})

	t.Run("not found", func(t *testing.T) {
		err := repository.DeleteCompany(context.Background(), uuid.New())
		req.ErrorContains(err, ErrNotFound.Error())
	})

}

func TestReadCompanyByID(t *testing.T) {
	req := require.New(t)
	err := repository.CreateCompany(context.Background(), companiesList[5])
	if err != nil {
		return
	}

	t.Run("read company success", func(t *testing.T) {
		company, err := repository.ReadCompanyByID(context.Background(), companiesList[5].ID())
		req.Empty(err)
		req.Equal(company.ID(), companiesList[5].ID())
	})

	t.Run("not found", func(t *testing.T) {
		_, err := repository.ReadCompanyByID(context.Background(), uuid.New())
		req.ErrorContains(err, ErrEmptyResult.Error())
	})

}

func createTestCompany() *company.Company {
	companyID := uuid.New()
	companyName, _ := name.NewName("ООО \"Лидген\"")
	companyInn, _ := inn.NewInn(1234567890)
	companyAddress, _ := address.NewAdress("г. Майкоп, ул. Краснооктябрьская 10А")
	testTime := time.Now()
	testCompany, _ := company.NewWithID(companyID, *companyName, *companyInn, *companyAddress, testTime, testTime)
	return testCompany
}
