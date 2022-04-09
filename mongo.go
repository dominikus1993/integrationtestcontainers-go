package integrationtestcontainers

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
)

type MongoContainerConfiguration struct {
	Image       string
	Port        int
	ExposedPort int
}

var DefaultMongoContainerConfiguration = &MongoContainerConfiguration{
	Image:       "mongo:5",
	Port:        27017,
	ExposedPort: 27017,
}

type MongoContainer struct {
	testcontainers.Container
	ConnectionString string
}

func NewMongoDbContainer(ctx context.Context, config *MongoContainerConfiguration) (*MongoContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        getValueOrDefault(config.Image, DefaultMongoContainerConfiguration.Image),
		ExposedPorts: []string{fmt.Sprintf("%d/tcp", config.ExposedPort)},
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
	mongoConnection := fmt.Sprintf("mongodb://%s:%s", host, port.Port())
	return &MongoContainer{mongoC, mongoConnection}, nil
}
