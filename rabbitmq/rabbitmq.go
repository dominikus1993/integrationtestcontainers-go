package rabbitmq

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
)

//var sql = []string{"/opt/mssql-tools/bin/sqlcmd", "-b", "-r", "1", "-U", config.username, "-P", config.password, "-i"}

type rabbitmqContainer struct {
	testcontainers.Container
	config *RabbitMqContainerConfiguration
}

func StartContainer(ctx context.Context, config *RabbitMqContainerConfiguration) (*rabbitmqContainer, error) {
	container, err := rabbitmq.Run(ctx,
		config.image,
		rabbitmq.WithAdminUsername(config.username),
		rabbitmq.WithAdminPassword(config.password),
	)
	if err != nil {
		return nil, err
	}
	return &rabbitmqContainer{container, config}, nil
}

func (container *rabbitmqContainer) ConnectionString(ctx context.Context) (string, error) {
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
