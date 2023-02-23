package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestRedisContainerPing(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	config := NewRedisContainerConfigurationBuilder().Build()
	// Arrange
	ctx := context.Background()

	container, err := StartContainer(ctx, config)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	connectionString, err := container.Url(ctx)
	if err != nil {
		t.Fatal(fmt.Errorf("can't download redis conectionstring, %w", err))
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	t.Cleanup(func() {
		if err := rdb.Close(); err != nil {
			t.Fatalf("failed to disconnect redis: %s", err)
		}
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		t.Fatal(fmt.Errorf("error when tring ping redis server: %w", err))
	}
}
