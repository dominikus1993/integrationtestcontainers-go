package mongodb

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	"github.com/dominikus1993/integrationtestcontainers-go/common"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type mongoContainer struct {
	testcontainers.Container
	config *MongoContainerConfiguration
}

func (container *mongoContainer) ConnectionString(ctx context.Context) (string, error) {
	host, err := container.Host(ctx)
	if err != nil {
		return "", err
	}
	port, err := container.MappedPort(ctx, nat.Port(fmt.Sprint(container.config.port)))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("mongodb://%s:%s", host, port.Port()), nil
}

func StartContainer(ctx context.Context, config *MongoContainerConfiguration) (*mongoContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        common.GetValueOrDefault(config.image, defaultMongoContainerConfiguration.image),
		ExposedPorts: []string{fmt.Sprintf("%d/tcp", config.exposedPort)},
		WaitingFor: wait.ForAll(
			wait.ForLog("Waiting for connections"),
			wait.ForListeningPort("27017/tcp"),
		),
	}
	if config.username != "" && config.password != "" {
		req.Env = map[string]string{"MONGO_INITDB_ROOT_USERNAME": config.username, "MONGO_INITDB_ROOT_PASSWORD": config.password}
	}

	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	return &mongoContainer{mongoC, config}, nil
}
