package integrationtestcontainers

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type RedisContainerConfiguration struct {
	Image       string
	Port        int
	ExposedPort int
}

var DefaultRedisContainerConfiguration = &RedisContainerConfiguration{
	Image:       "redis:6",
	Port:        6379,
	ExposedPort: 6379,
}

type RedisContainer struct {
	testcontainers.Container
	ConnectionString string
}

func StartRedisContainer(ctx context.Context, config *RedisContainerConfiguration) (*RedisContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        getValueOrDefault(config.Image, DefaultMongoContainerConfiguration.Image),
		ExposedPorts: []string{fmt.Sprintf("%d/tcp", config.ExposedPort)},
		WaitingFor:   wait.ForLog("* Ready to accept connections"),
	}
	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	host, err := mongoC.Host(ctx)
	if err != nil {
		return nil, err
	}
	port, err := mongoC.MappedPort(ctx, nat.Port(fmt.Sprint(config.Port)))
	if err != nil {
		return nil, err
	}
	mongoConnection := fmt.Sprintf("redis://%s:%s", host, port.Port())
	return &RedisContainer{mongoC, mongoConnection}, nil
}
