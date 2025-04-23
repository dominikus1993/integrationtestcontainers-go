package sqlserver

import (
	"context"
	"fmt"
	"testing"

	"database/sql"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/stretchr/testify/assert"
)

func TestSqlServerContainerPing(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	config := NewSqlServerContainerConfigurationBuilder().Build()
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
		t.Fatal(fmt.Errorf("can't download sql conectionstring, %w", err))
	}
	client, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		t.Fatal(fmt.Errorf("error creating sql client: %w", err))
	}

	t.Cleanup(func() {
		if err := client.Close(); err != nil {
			t.Fatalf("failed to disconnect container: %s", err)
		}
	})

	subject := client.QueryRow("SELECT 1")
	assert.NoError(t, subject.Err())
}
