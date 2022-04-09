package integrationtestcontainers

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainerConfiguration struct {
	Image       string
	Port        int
	ExposedPort int
	Database    string
	Username    string
	Password    string
}

var DefaultPostgresContainerConfiguration = &PostgresContainerConfiguration{
	Image:       "postgres:14-bullseye",
	Port:        5432,
	ExposedPort: 5432,
	Database:    "postgres",
	Username:    "postgres",
	Password:    "postgres",
}

type PostgresContainer struct {
	testcontainers.Container
	ConnectionString string
}

func NewPostgreSqlContainer(ctx context.Context, config *PostgresContainerConfiguration) (*PostgresContainer, error) {
	var port string = getValueOrDefault(fmt.Sprint(config.Port), "5432")
	req := testcontainers.ContainerRequest{
		Image:        getValueOrDefault(config.Image, DefaultPostgresContainerConfiguration.Image),
		ExposedPorts: []string{fmt.Sprintf("%d/tcp", config.ExposedPort)},
		Env:          map[string]string{"POSTGRES_USER": config.Username, "POSTGRES_PASSWORD": config.Password, "POSTGRES_DB": config.Database},
		WaitingFor:   wait.NewExecStrategy([]string{"pg_isready", "--host", "localhost", "--port", port}),
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
	cport, err := mongoC.MappedPort(ctx, nat.Port(port))
	if err != nil {
		return nil, err
	}
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.Username, config.Password, host, cport.Port(), config.Database)
	return &PostgresContainer{mongoC, connectionString}, nil
}
