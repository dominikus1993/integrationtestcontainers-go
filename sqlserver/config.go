package sqlserver

import "github.com/dominikus1993/integrationtestcontainers-go/common"

var defaultSqlServerContainerConfiguration = &SqlServerContainerConfiguration{
	image:       "mcr.microsoft.com/mssql/server:2019-CU18-ubuntu-20.04",
	port:        1433,
	exposedPort: 1433,
	database:    "master",
	username:    "sa",
	password:    "jp2GMD(!)xDDD",
}

type SqlServerContainerConfiguration struct {
	image       string
	port        int
	exposedPort int
	database    string
	username    string
	password    string
}

type sqlServerContainerConfigurationBuilder struct {
	image       string
	port        int
	exposedPort int
	database    string
	username    string
	password    string
}

func NewSqlServerContainerConfigurationBuilder() *sqlServerContainerConfigurationBuilder {
	return &sqlServerContainerConfigurationBuilder{
		image:       defaultSqlServerContainerConfiguration.image,
		port:        defaultSqlServerContainerConfiguration.port,
		exposedPort: defaultSqlServerContainerConfiguration.exposedPort,
		database:    defaultSqlServerContainerConfiguration.database,
		username:    defaultSqlServerContainerConfiguration.username,
		password:    defaultSqlServerContainerConfiguration.password,
	}
}

func (builder *sqlServerContainerConfigurationBuilder) WithPort(port int) common.Builder[SqlServerContainerConfiguration] {
	builder.port = port
	return builder
}

func (builder *sqlServerContainerConfigurationBuilder) WithImage(image string) common.Builder[SqlServerContainerConfiguration] {
	builder.image = image
	return builder
}

func (builder *sqlServerContainerConfigurationBuilder) WithUsername(username string) common.Builder[SqlServerContainerConfiguration] {
	builder.username = username
	return builder
}

func (builder *sqlServerContainerConfigurationBuilder) WithPassword(password string) common.Builder[SqlServerContainerConfiguration] {
	builder.password = password
	return builder
}

func (builder *sqlServerContainerConfigurationBuilder) WithDatabase(database string) common.Builder[SqlServerContainerConfiguration] {
	builder.database = database
	return builder
}

func (builder *sqlServerContainerConfigurationBuilder) Build() *SqlServerContainerConfiguration {
	return &SqlServerContainerConfiguration{
		image:       builder.image,
		port:        builder.port,
		exposedPort: builder.exposedPort,
		database:    builder.database,
		username:    builder.username,
		password:    builder.password,
	}
}
