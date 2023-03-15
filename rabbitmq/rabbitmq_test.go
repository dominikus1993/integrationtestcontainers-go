package rabbitmq

import (
	"context"
	"fmt"
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
)

func TestRabbitMqContainerPing(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	config := NewRabbitMqContainerConfigurationBuilder().Build()
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

	url, err := container.ConnectionString(ctx)
	if err != nil {
		t.Fatal(fmt.Errorf("can't download rabbit conectionstring, %w", err))
	}
	conn, err := amqp.Dial(url)
	if err != nil {
		t.Fatal(fmt.Errorf("can't connect to rabbitmq, %w", err))
	}

	t.Cleanup(func() {
		if err := conn.Close(); err != nil {
			t.Fatalf("failed to terminate rabbitmq connection: %s", err)
		}
	})

	closed := conn.IsClosed()
	assert.False(t, closed)
}
