package redis

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type redisContainer struct {
	testcontainers.Container
	config *RedisContainerConfiguration
}

func StartContainer(ctx context.Context, config *RedisContainerConfiguration) (*redisContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        config.image,
		ExposedPorts: []string{fmt.Sprintf("%d/tcp", config.exposedPort)},
		WaitingFor:   wait.ForLog("* Ready to accept connections"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	return &redisContainer{container, config}, nil
}

func (container *redisContainer) Url(ctx context.Context) (string, error) {
	mappedPort, err := container.MappedPort(ctx, nat.Port(fmt.Sprint(container.config.port)))
	if err != nil {
		return "", err
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return "", err
	}

	uri := fmt.Sprintf("redis://%s:%s", hostIP, mappedPort.Port())

	return uri, nil
}
