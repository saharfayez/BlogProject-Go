package utils

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var ShutDownTestContainer func()

func StartTestContainer(ctx context.Context) (string, error) {
	var env = map[string]string{
		"POSTGRES_PASSWORD": "postgres",
		"POSTGRES_USER":     "postgres",
		"POSTGRES_DB":       "postgres",
	}
	var port = "5432/tcp"
	// Define the PostgreSQL container request
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15-alpine",
		ExposedPorts: []string{port},
		Env:          env,
		WaitingFor:   wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
	}

	// Start the PostgreSQL container
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return "", fmt.Errorf("Failed to start PostgreSQL container: %v", err)
	}

	// Get the host and port of the PostgreSQL container
	_, err = container.Host(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get PostgreSQL container host: %v", err)
	}

	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return "", fmt.Errorf("failed to get PostgreSQL container port: %v", err)
	}
	fmt.Printf("PORT NUMBER: %d", mappedPort.Port())

	//time.Sleep(5 * time.Second)
	// Construct the connection string for GORM

	dsn := fmt.Sprintf("host=127.0.0.1 user=postgres password=postgres dbname=postgres port=%s sslmode=disable TimeZone=Asia/Jakarta",
		mappedPort.Port())

	ShutDownTestContainer = func() {
		if err := container.Terminate(ctx); err != nil {
			fmt.Printf("Failed to terminate PostgreSQL container: %v\n", err)
		}
	}
	return dsn, nil
}
