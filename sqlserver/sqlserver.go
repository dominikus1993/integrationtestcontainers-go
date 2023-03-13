package sqlserver

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
	config *SqlServerContainerConfiguration
}

func StartContainer(ctx context.Context, config *SqlServerContainerConfiguration) (*msSqlContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        config.image,
		ExposedPorts: []string{fmt.Sprintf("%d/tcp", config.exposedPort)},
		Env:          map[string]string{"SQLCMDUSER": config.username, "SQLCMDDBNAME": config.database, "MSSQL_SA_PASSWORD": config.password, "SQLCMDPASSWORD": config.password, "ACCEPT_EULA": "Y"},
		WaitingFor:   wait.ForExec([]string{"/opt/mssql-tools/bin/sqlcmd", "-Q", "SELECT 1;"}),
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
	uri := fmt.Sprintf("sqlserver://%s:%s@%s?port=%s", container.config.username, container.config.password, hostIP, port)

	return uri, nil
}
