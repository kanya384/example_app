package integration_test

import (
	"companies/internal/config"
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {

}

func init_() {
	ctx := context.Background()

	cfg, err := config.InitConfig("COMPANIES")
	if err != nil {
		panic(err)
	}
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}

	list, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range list {
		cli.ContainerStop(ctx, container.ID, nil)
	}

	postgresContainerID, err := startPostgresContainer(ctx, cli, cfg)
	if err != nil {
		panic(fmt.Errorf("error starting postgres container: %w", err))
	}

	err = removeContainers(ctx, cli, postgresContainerID)
	if err != nil {
		panic(fmt.Errorf("error removing containers, please remove them manualy: %w", err))
	}
}

func startPostgresContainer(ctx context.Context, cli *client.Client, cfg *config.Config) (continerID string, err error) {
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        "postgres",
		ExposedPorts: nat.PortSet{"5432": struct{}{}},
		Env: []string{
			fmt.Sprintf("POSTGRES_USER=%s", cfg.PG.User),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", cfg.PG.Pass),
			fmt.Sprintf("POSTGRES_DB=%s", cfg.PG.DbName),
		},
	}, &container.HostConfig{
		PortBindings: map[nat.Port][]nat.PortBinding{nat.Port("5432"): {{HostIP: "127.0.0.1", HostPort: "5432"}}},
	}, nil, nil, "postgres_test")
	if err != nil {
		return
	}
	continerID = resp.ID

	if err := cli.ContainerStart(ctx, continerID, types.ContainerStartOptions{}); err != nil {
		return continerID, err
	}

	return
}

func removeContainers(ctx context.Context, cli *client.Client, containers ...string) (err error) {
	for _, container := range containers {
		err = cli.ContainerStop(ctx, container, nil)
		if err != nil {
			return
		}
		err = cli.ContainerRemove(ctx, container, types.ContainerRemoveOptions{})
		if err != nil {
			return
		}
	}
	return
}
