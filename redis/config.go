package redis

import "github.com/dominikus1993/integrationtestcontainers-go/common"

var defaultMongoContainerConfiguration = &RedisContainerConfiguration{
	image:       "redis:6",
	port:        6379,
	exposedPort: 6379,
}

type RedisContainerConfiguration struct {
	image       string
	port        int
	exposedPort int
}

type redisContainerConfigurationBuilder struct {
	image       string
	port        int
	exposedPort int
}

func NewRedisContainerConfigurationBuilder() common.Builder[RedisContainerConfiguration] {
	return &redisContainerConfigurationBuilder{
		image:       defaultMongoContainerConfiguration.image,
		port:        defaultMongoContainerConfiguration.port,
		exposedPort: defaultMongoContainerConfiguration.exposedPort,
	}
}

func (builder *redisContainerConfigurationBuilder) WithPort(port int) common.Builder[RedisContainerConfiguration] {
	builder.port = port
	return builder
}

func (builder *redisContainerConfigurationBuilder) WithImage(image string) common.Builder[RedisContainerConfiguration] {
	builder.image = image
	return builder
}

func (builder *redisContainerConfigurationBuilder) WithUsername(username string) common.Builder[RedisContainerConfiguration] {
	return builder
}

func (builder *redisContainerConfigurationBuilder) WithPassword(password string) common.Builder[RedisContainerConfiguration] {
	return builder
}

func (builder *redisContainerConfigurationBuilder) Build() *RedisContainerConfiguration {
	return &RedisContainerConfiguration{
		image:       builder.image,
		port:        builder.port,
		exposedPort: builder.exposedPort,
	}
}
