package test_utils

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/go-faster/errors"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
)

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/ryota0624/go-dev-with-ent/migrations"
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

func PSQLContainer(ctx context.Context) (testcontainers.Container, error) {
	psqlPort, err := nat.NewPort("tcp", "5432")
	if err != nil {
		return nil, errors.Wrap(err, "failed to NewPort")
	}

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{fmt.Sprintf("%s/%s", psqlPort.Port(), psqlPort.Proto())},
		WaitingFor: wait.ForSQL(psqlPort, "postgres", func(port nat.Port) string {
			return fmt.Sprintf("postgres://localhost:%s?user=postgres&password=password&sslmode=disable", port.Port())
		}),
		Env: map[string]string{
			"POSTGRES_PASSWORD": "password",
		},
	}

	psqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to GenericContainer creation")
	}

	return psqlContainer, nil
}

func StartRDB(ctx context.Context) (*RDBContainer, error) {
	psqlPort, err := nat.NewPort("tcp", "5432")
	if err != nil {
		return nil, errors.Wrap(err, "failed to NewPort")
	}
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{fmt.Sprintf("%s/%s", psqlPort.Port(), psqlPort.Proto())},
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

	d, err := iofs.New(migrations.MigrationFiles, ".")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, fmt.Sprintf("%s/sample?user=postgres&password=passward&sslmode=disable", psqlEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to migrate.New")
	}

	if err := m.Up(); err != nil {
		return nil, errors.Wrap(err, "failed to migration Up")
	}

	return &RDBContainer{psqlContainer}, nil
}
