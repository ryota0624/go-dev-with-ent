package main

import (
	"context"
	"log"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
	"github.com/ryota0624/go-dev-with-ent/ent"
)

func main() {
	// atlas.sumがない場合は atlas migrate hash --force で作る
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=passward dbname=sample sslmode=disable")
	if err != nil {
		log.Fatalf("failed connecting to postgreql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Create a local migration directory.
	dir, err := migrate.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Write migration diff.
	err = client.Schema.Diff(ctx, schema.WithDir(dir), schema.WithSumFile())
	// You can use the following method to give the migration files a name.
	// err = client.Schema.NamedDiff(ctx, "migration_name", schema.WithDir(dir))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
