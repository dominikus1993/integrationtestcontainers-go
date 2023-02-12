package mongodb

import "github.com/dominikus1993/integrationtestcontainers-go/common"

var defaultMongoContainerConfiguration = &MongoContainerConfiguration{
	image:       "mongo:6",
	port:        27017,
	exposedPort: 27017,
	username:    "",
	password:    "",
}

type MongoContainerConfiguration struct {
	image       string
	port        int
	exposedPort int
	username    string
	password    string
}

type mongoContainerConfigurationBuilder struct {
	image       string
	port        int
	exposedPort int
	username    string
	password    string
}

func NewMongoContainerConfigurationBuilder() common.Builder[MongoContainerConfiguration] {
	return &mongoContainerConfigurationBuilder{
		image:       defaultMongoContainerConfiguration.image,
		port:        defaultMongoContainerConfiguration.port,
		exposedPort: defaultMongoContainerConfiguration.exposedPort,
		username:    defaultMongoContainerConfiguration.username,
		password:    defaultMongoContainerConfiguration.password,
	}
}

func (builder *mongoContainerConfigurationBuilder) WithPort(port int) common.Builder[MongoContainerConfiguration] {
	builder.port = port
	return builder
}

func (builder *mongoContainerConfigurationBuilder) WithImage(image string) common.Builder[MongoContainerConfiguration] {
	builder.image = image
	return builder
}

func (builder *mongoContainerConfigurationBuilder) WithUsername(username string) common.Builder[MongoContainerConfiguration] {
	builder.username = username
	return builder
}

func (builder *mongoContainerConfigurationBuilder) WithPassword(password string) common.Builder[MongoContainerConfiguration] {
	builder.password = password
	return builder
}

func (builder *mongoContainerConfigurationBuilder) Build() *MongoContainerConfiguration {
	return &MongoContainerConfiguration{
		image:       builder.image,
		port:        builder.port,
		exposedPort: builder.exposedPort,
		username:    builder.username,
		password:    builder.password,
	}
}
