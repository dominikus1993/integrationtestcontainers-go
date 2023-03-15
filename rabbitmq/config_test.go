package rabbitmq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkToSlice-4       1302279               897.8 ns/op           688 B/op         12 allocs/op
func TestSqlServerConfigBuilder(t *testing.T) {
	builder := NewRabbitMqContainerConfigurationBuilder()

	config := builder.WithImage("mongo:6").WithPort(27017).WithPassword("kremowki").WithUsername("janpawel2").Build()
	assert.NotNil(t, config)
	assert.Equal(t, "mongo:6", config.image)
	assert.Equal(t, 27017, config.port)
	assert.Equal(t, "kremowki", config.password)
	assert.Equal(t, "janpawel2", config.username)
}

func BenchmarkToSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		builder := NewRabbitMqContainerConfigurationBuilder()

		_ = builder.WithImage("mongo:6").WithPort(27017).WithPassword("kremowki").WithUsername("janpawel2").Build()
	}
}
