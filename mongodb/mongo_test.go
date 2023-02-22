package mongodb

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoContainerPing(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	config := NewMongoContainerConfigurationBuilder().Build()
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

	connectionString, err := container.ConnectionString(ctx)
	if err != nil {
		t.Fatal(fmt.Errorf("can't download mongo conectionstring, %w", err))
	}
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		t.Fatal(fmt.Errorf("error creating mongo client: %w", err))
	}

	err = mongoClient.Connect(ctx)
	if err != nil {
		t.Fatal(fmt.Errorf("error connectiong mongo client: %w", err))
	}

	t.Cleanup(func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			t.Fatalf("failed to disconnect container: %s", err)
		}
	})

	subject := mongoClient.Ping(ctx, nil)
	assert.NoError(t, subject)
}
