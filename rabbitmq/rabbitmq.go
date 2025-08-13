package rabbitmq

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

//var sql = []string{"/opt/mssql-tools/bin/sqlcmd", "-b", "-r", "1", "-U", config.username, "-P", config.password, "-i"}

type msSqlContainer struct {
	testcontainers.Container
	config *RabbitMqContainerConfiguration
}

func StartContainer(ctx context.Context, config *RabbitMqContainerConfiguration) (*msSqlContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        config.image,
		ExposedPorts: []string{fmt.Sprintf("%d/tcp", config.exposedPort)},
		Env:          map[string]string{"RABBITMQ_DEFAULT_USER": config.username, "RABBITMQ_DEFAULT_PASS": config.password},
		WaitingFor: wait.ForAll(
			wait.ForLog("Server startup complete"),
		),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	return &msSqlContainer{container, config}, nil
}

func (container *msSqlContainer) ConnectionString(ctx context.Context) (string, error) {
	mappedPort, err := container.MappedPort(ctx, nat.Port(fmt.Sprint(container.config.port)))
	if err != nil {
		return "", err
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return "", err
	}

	port := mappedPort.Port()
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s", container.config.username, container.config.password, hostIP, port)

	return uri, nil
}
