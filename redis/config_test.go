package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkToSlice-4       1302279               897.8 ns/op           688 B/op         12 allocs/op
func TestRedisConfigBuilder(t *testing.T) {
	builder := NewRedisContainerConfigurationBuilder()

	config := builder.WithImage("redis:6").WithPort(6379).Build()
	assert.NotNil(t, config)
	assert.Equal(t, "redis:6", config.image)
	assert.Equal(t, 6379, config.port)
}

func BenchmarkToSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		builder := NewRedisContainerConfigurationBuilder()

		_ = builder.WithImage("redis:6").WithPort(6379).WithPassword("kremowki").WithUsername("janpawel2").Build()
	}
}
