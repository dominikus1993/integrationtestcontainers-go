package sqlserver

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mssql"
)

//var sql = []string{"/opt/mssql-tools/bin/sqlcmd", "-b", "-r", "1", "-U", config.username, "-P", config.password, "-i"}

type msSqlContainer struct {
	testcontainers.Container
	config *SqlServerContainerConfiguration
}

func StartContainer(ctx context.Context, config *SqlServerContainerConfiguration) (*msSqlContainer, error) {
	mssqlContainer, err := mssql.Run(ctx,
		config.image,
		mssql.WithAcceptEULA(),
		mssql.WithPassword(config.password),
	)
	if err != nil {
		return nil, err
	}
	return &msSqlContainer{mssqlContainer, config}, nil
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

func (container *msSqlContainer) ExecCommand(ctx context.Context, scriptContent string) (string, error) {
	scriptFilePath := strings.Join([]string{"./tmp", fmt.Sprintf("%s.sql", uuid.NewString())}, "/")
	err := container.CopyToContainer(ctx, []byte(scriptContent), scriptFilePath, 493)
	if err != nil {
		return "", err
	}
	_, reader, err := container.Exec(ctx, []string{"/opt/mssql-tools/bin/sqlcmd", "-b", "-r", "1", "-U", container.config.username, "-P", container.config.password, "-i", scriptFilePath})
	if err != nil {
		return "", err
	}
	if b, err := io.ReadAll(reader); err == nil {
		return string(b), nil
	}
	return "", err
}
