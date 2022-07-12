package user

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/google/uuid"
	"github.com/ryota0624/go-dev-with-ent/ent"
	"github.com/ryota0624/go-dev-with-ent/model/test_utils"
	"github.com/stretchr/testify/assert"
	"log"

	"testing"
)

func Test_PSQL_Query(t *testing.T) {
	t.Run("run select 1", func(t *testing.T) {
		ctx := context.Background()
		rdbContainer, err := test_utils.PSQLContainer(ctx)
		if err != nil {
			t.Fatalf("failed to PSQLContainer creation: %v", err)
		}
		psqlEndpoint, err := rdbContainer.Endpoint(context.Background(), "postgres")
		if err != nil {
			t.Fatalf("failed to get postgres endpoint: %v", err)
		}
		t.Logf("endpoint: %s", psqlEndpoint)
		source := fmt.Sprintf("%s/sample?user=postgres&password=password&sslmode=disable", psqlEndpoint)

		db, err := sql.Open("postgres", source)
		if err != nil {
			t.Fatalf("failed to sql.Open: %v", err)
		}

		result, err := db.Query("select 1")
		if err != nil {
			log.Fatalf("main sql.Open error err: %v", err)
		}

		defer result.Close()
		for result.Next() {
			var number int
			if err := result.Scan(&number); err != nil {
				t.Fatalf("failed to Scan Result: %v", err)
			}
			t.Logf("number=%d", number)
		}

		if err := result.Err(); err != nil {
			t.Fatalf("result.Err()=$%v", err)
		}

		defer db.Close()
		if err != nil {
			t.Fatalf("failed to start RDB %v", err)
		}
		defer rdbContainer.Terminate(ctx)
	})
}

func Test_factoryImpl_Create(t *testing.T) {

	t.Run("Passed Group Should Exists", func(t *testing.T) {
		ctx := context.Background()
		rdbContainer, err := test_utils.StartRDB(ctx)
		if err != nil {
			t.Fatalf("failed to start RDB %v", err)
		}
		defer rdbContainer.Terminate(ctx)
		client, err := ent.Open(dialect.Postgres, rdbContainer.DatabaseEndpoint())
		if err != nil {
			log.Fatal(err)
		}
		ctx = WithEntClient(ctx, client)
		factory := &factoryImpl{}
		_, err = factory.Create(ctx, uuid.New(), 10, "some")
		assert.ErrorIs(t, err, ErrGroupIdShouldExists)
	})

	t.Run("Create Should Pass if Group Exists", func(t *testing.T) {
		ctx := context.Background()
		rdbContainer, err := test_utils.StartRDB(ctx)
		if err != nil {
			t.Fatalf("failed to start RDB %v", err)
		}
		defer rdbContainer.Terminate(ctx)
		client, err := ent.Open(dialect.Postgres, rdbContainer.DatabaseEndpoint())
		if err != nil {
			log.Fatal(err)
		}

		groupId := uuid.New()
		err = client.Group.Create().SetID(groupId).SetName(groupId.String()).Exec(ctx)
		if err != nil {
			log.Fatal(err)
		}

		ctx = WithEntClient(ctx, client)
		factory := &factoryImpl{}
		_, err = factory.Create(ctx, groupId, 10, "some")
		assert.Nil(t, err)
	})
}
