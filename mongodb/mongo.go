package mongodb

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	"github.com/dominikus1993/integrationtestcontainers-go/common"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MongoContainer struct {
	testcontainers.Container
	config *MongoContainerConfiguration
}

func (container *MongoContainer) ConnectionString(ctx context.Context) (string, error) {
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

func StartMongoDbContainer(ctx context.Context, config *MongoContainerConfiguration) (*MongoContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        common.GetValueOrDefault(config.image, defaultMongoContainerConfiguration.image),
		ExposedPorts: []string{fmt.Sprintf("%d/tcp", config.exposedPort)},
		WaitingFor:   wait.NewExecStrategy([]string{"mongo", fmt.Sprintf("localhost:%d", config.port), "--eval", "db.runCommand(\"ping\").ok", "--quiet"}),
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
	return &MongoContainer{mongoC, config}, nil
}
