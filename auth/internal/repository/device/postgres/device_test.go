package postgres

import (
	"auth/internal/domain/device"
	"auth/internal/domain/device/agent"
	"auth/internal/domain/device/deviceID"
	"auth/internal/domain/device/ip"
	"auth/internal/domain/device/refreshToken"
	userDomain "auth/internal/domain/user"
	"auth/internal/domain/user/email"
	"auth/internal/domain/user/name"
	"auth/internal/domain/user/pass"
	"auth/internal/domain/user/phone"
	"auth/internal/domain/user/surname"
	user "auth/internal/repository/user/postgres"

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
	testUser            *userDomain.User
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

	usersRepo, err := user.New(pg, user.Options{})
	if err != nil {
		log.Fatalf("error initializing users repository: %s", err.Error())
	}

	testUser = createTestUser()
	err = usersRepo.CreateUser(context.Background(), testUser)
	if err != nil {
		log.Fatalf("error creating user for testing purposes: %s", err.Error())
	}

	os.Exit(m.Run())
}

func TestCreateDevice(t *testing.T) {
	req := require.New(t)
	device := createTestDevice()

	fmt.Println(device)

	t.Run("create device success", func(t *testing.T) {
		err := repository.CreateDevice(context.Background(), device)
		req.Empty(err)
	})

	t.Run("duplicate key error", func(t *testing.T) {
		err := repository.CreateDevice(context.Background(), device)
		req.ErrorContains(err, ErrDuplicateKey.Error())
	})

}

func TestUpdateDevice(t *testing.T) {
	req := require.New(t)
	deviceForUpdate := createTestDevice()

	err := repository.CreateDevice(context.Background(), deviceForUpdate)
	if err != nil {
		return
	}

	newDevice, _ := device.NewWithID(deviceForUpdate.ID(), deviceForUpdate.CreatedAt(), time.Now(), deviceForUpdate.UserID(), deviceID.DeviceID("1111111111111111111-11"), ip.Ip("10.2.3.4"), agent.Agent("adnroid galaxy s5"), device.Android, refreshToken.RefreshToken("test"), time.Now(), time.Now())

	updateFunc := func(oldDevice *device.Device) (*device.Device, error) {
		return device.NewWithID(oldDevice.ID(), oldDevice.CreatedAt(), newDevice.ModifiedAt(), oldDevice.UserID(), newDevice.DeviceID(), newDevice.Ip(), newDevice.Agent(), newDevice.Type(), newDevice.RefreshToken(), newDevice.RefreshExpiration(), newDevice.LastSeen())
	}

	t.Run("update device success", func(t *testing.T) {

		res, err := repository.UpdateDevice(context.Background(), deviceForUpdate.ID(), updateFunc)
		req.Empty(err)
		req.Equal(res.ID(), newDevice.ID())
		req.Equal(res.UserID(), newDevice.UserID())
		req.Equal(res.DeviceID(), newDevice.DeviceID())
		req.Equal(res.Ip(), newDevice.Ip())
		req.Equal(res.Agent(), newDevice.Agent())
		req.Equal(res.Type(), newDevice.Type())
		req.Equal(res.RefreshToken(), newDevice.RefreshToken())
		req.Equal(res.RefreshExpiration(), newDevice.RefreshExpiration())
		req.Equal(res.LastSeen(), newDevice.LastSeen())
		req.NotEmpty(res.CreatedAt())
	})

	t.Run("not found", func(t *testing.T) {
		_, err := repository.UpdateDevice(context.Background(), uuid.New(), updateFunc)
		req.ErrorContains(err, ErrEmptyResult.Error())
	})
}

func TestDeleteDevice(t *testing.T) {
	req := require.New(t)
	device := createTestDevice()
	err := repository.CreateDevice(context.Background(), device)
	if err != nil {
		log.Fatalf("error creating test device: %s", err.Error())
	}

	t.Run("delete device success", func(t *testing.T) {
		err := repository.DeleteDevice(context.Background(), uuid.UUID(device.ID()))
		req.Empty(err)
	})

	t.Run("not found", func(t *testing.T) {
		err := repository.DeleteDevice(context.Background(), uuid.UUID(device.ID()))
		req.ErrorContains(err, ErrNotFound.Error())
	})

}

func TestReadDeviceByID(t *testing.T) {
	req := require.New(t)
	device := createTestDevice()
	err := repository.CreateDevice(context.Background(), device)
	if err != nil {
		log.Fatalf("error creating test device: %s", err.Error())
	}

	t.Run("success", func(t *testing.T) {
		deviceR, err := repository.ReadDeviceByID(context.Background(), uuid.UUID(device.ID()))
		req.Empty(err)
		req.Equal(device.ID(), deviceR.ID())
	})

	t.Run("not found", func(t *testing.T) {
		_, err := repository.ReadDeviceByID(context.Background(), uuid.New())
		req.ErrorContains(err, ErrEmptyResult.Error())
	})

}

func createTestDevice() *device.Device {
	userID := testUser.ID()
	deviceID, _ := deviceID.NewDeviceID("000000000000000-00")
	ip, _ := ip.NewIp("10.2.0.1")
	agent, _ := agent.NewAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063")
	dtype := device.Web
	refreshToken, _ := refreshToken.New()
	refreshExp := time.Now()
	lastSeen := time.Now()
	testDevice, _ := device.New(userID, *deviceID, *ip, *agent, dtype, *refreshToken, refreshExp, lastSeen)
	return testDevice
}

func createTestUser() *userDomain.User {
	name, _ := name.NewName("user")
	surname, _ := surname.NewSurname("surname")
	phone, _ := phone.NewPhone("+7 (900) 000-00-00")
	email, _ := email.NewEmail("test01@mail.ru")
	pass := pass.Pass("Password123!")
	role := userDomain.Administrator
	testUser, _ := userDomain.New(*name, *surname, *phone, pass, *email, role)
	return testUser
}
