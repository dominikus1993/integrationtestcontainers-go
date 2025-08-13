package rabbitmq

import "github.com/dominikus1993/integrationtestcontainers-go/common"

var defaultRabbitMqContainerConfiguration = &RabbitMqContainerConfiguration{
	image:       "rabbitmq:3.12.11-management-alpine",
	port:        5672,
	exposedPort: 5672,
	username:    "guest",
	password:    "jp2GMD(!)xDDD",
}

type RabbitMqContainerConfiguration struct {
	image       string
	port        int
	exposedPort int
	username    string
	password    string
}

type RabbitMqContainerConfigurationBuilder struct {
	image       string
	port        int
	exposedPort int
	username    string
	password    string
}

func NewRabbitMqContainerConfigurationBuilder() *RabbitMqContainerConfigurationBuilder {
	return &RabbitMqContainerConfigurationBuilder{
		image:       defaultRabbitMqContainerConfiguration.image,
		port:        defaultRabbitMqContainerConfiguration.port,
		exposedPort: defaultRabbitMqContainerConfiguration.exposedPort,
		username:    defaultRabbitMqContainerConfiguration.username,
		password:    defaultRabbitMqContainerConfiguration.password,
	}
}

func (builder *RabbitMqContainerConfigurationBuilder) WithPort(port int) common.Builder[RabbitMqContainerConfiguration] {
	builder.port = port
	return builder
}

func (builder *RabbitMqContainerConfigurationBuilder) WithImage(image string) common.Builder[RabbitMqContainerConfiguration] {
	builder.image = image
	return builder
}

func (builder *RabbitMqContainerConfigurationBuilder) WithUsername(username string) common.Builder[RabbitMqContainerConfiguration] {
	builder.username = username
	return builder
}

func (builder *RabbitMqContainerConfigurationBuilder) WithPassword(password string) common.Builder[RabbitMqContainerConfiguration] {
	builder.password = password
	return builder
}

func (builder *RabbitMqContainerConfigurationBuilder) Build() *RabbitMqContainerConfiguration {
	return &RabbitMqContainerConfiguration{
		image:       builder.image,
		port:        builder.port,
		exposedPort: builder.exposedPort,
		username:    builder.username,
		password:    builder.password,
	}
}
