package integrationtestcontainers

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MongoContainerConfiguration struct {
	Image       string
	Port        int
	ExposedPort int
	Username    string
	Password    string
}

var DefaultMongoContainerConfiguration = &MongoContainerConfiguration{
	Image:       "mongo:5",
	Port:        27017,
	ExposedPort: 27017,
	Username:    "",
	Password:    "",
}

type MongoContainer struct {
	testcontainers.Container
	ConnectionString string
}

func StartMongoDbContainer(ctx context.Context, config *MongoContainerConfiguration) (*MongoContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        getValueOrDefault(config.Image, DefaultMongoContainerConfiguration.Image),
		ExposedPorts: []string{fmt.Sprintf("%d/tcp", config.ExposedPort)},
		WaitingFor:   wait.NewExecStrategy([]string{"mongo", fmt.Sprintf("localhost:%d", config.Port), "--eval", "db.runCommand(\"ping\").ok", "--quiet"}),
	}
	if config.Username != "" && config.Password != "" {
		req.Env = map[string]string{"MONGO_INITDB_ROOT_USERNAME": config.Username, "MONGO_INITDB_ROOT_PASSWORD": config.Password}
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
