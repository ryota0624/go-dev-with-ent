package model

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/go-faster/errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
)

type RDBContainer struct {
	testcontainers.Container
}

func (container *RDBContainer) DatabaseEndpoint() string {
	psqlEndpoint, err := container.Container.Endpoint(context.Background(), "postgres")
	if err != nil {
		panic(errors.Wrap(err, "psql container should expose endpoint"))
	}
	return fmt.Sprintf("%s/sample?user=postgres&password=passward&sslmode=disable", psqlEndpoint)
}

func StartRDB(ctx context.Context) (*RDBContainer, error) {
	psqlPort, err := nat.NewPort("tcp", "5432")
	if err != nil {
		return nil, errors.Wrap(err, "failed to NewPort")
	}
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor: wait.ForSQL(psqlPort, "postgres", func(port nat.Port) string {
			return fmt.Sprintf("postgres://localhost:%s?user=postgres&password=passward&sslmode=disable", port.Port())
		}),

		//WaitingFor:   wait.ForLog("database system is ready to accept connections").WithPollInterval(time.Second),
		Env: map[string]string{
			"POSTGRES_PASSWORD": "passward",
		},
	}

	psqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	psqlEndpoint, err := psqlContainer.Endpoint(ctx, "postgres")
	if err != nil {
		return nil, errors.Wrap(err, "failed Get Endpoint")
	}
	m, err := migrate.New(
		"file://../migrations",
		fmt.Sprintf("%s/sample?user=postgres&password=passward&sslmode=disable", psqlEndpoint),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to migrate.New")
	}

	if err := m.Up(); err != nil {
		return nil, errors.Wrap(err, "failed to migration Up")
	}

	return &RDBContainer{psqlContainer}, nil
}

func TestWithRDB(t *testing.T) {
	ctx := context.Background()
	psqlPort, err := nat.NewPort("tcp", "5432")
	if err != nil {
		t.Fatalf("failed to construct Port")
	}
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor: wait.ForSQL(psqlPort, "postgres", func(port nat.Port) string {
			return fmt.Sprintf("postgres://localhost:%s?user=postgres&password=passward&sslmode=disable", port.Port())
		}),

		//WaitingFor:   wait.ForLog("database system is ready to accept connections").WithPollInterval(time.Second),
		Env: map[string]string{
			"POSTGRES_PASSWORD": "passward",
		},
	}

	psqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}
	defer psqlContainer.Terminate(ctx)
	psqlEndpoint, err := psqlContainer.Endpoint(ctx, "postgres")
	if err != nil {
		t.Fatalf("failed Get Endpoint: %v", err)
	}
	t.Logf("psql Endpoint %s", psqlEndpoint)
	m, err := migrate.New(
		"file://../migrations",
		fmt.Sprintf("%s/sample?user=postgres&password=passward&sslmode=disable", psqlEndpoint),
	)
	if err != nil {
		t.Fatalf("failed New migration: %v", err)
	}

	if err := m.Up(); err != nil {
		t.Fatalf("failed to migration Up: %v", err)
	}
}

func TestPort(t *testing.T) {
	psqlPort, err := nat.NewPort("tcp", "5432")
	if err != nil {
		t.Fatalf("failed %v", err)
	}

	t.Logf("spec=%s", fmt.Sprintf("%s/%s", psqlPort.Port(), psqlPort.Proto()))

}
