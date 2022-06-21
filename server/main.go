package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ryota0624/go-dev-with-ent/ent"
	"github.com/ryota0624/go-dev-with-ent/ent/ogent"
	"github.com/ryota0624/go-dev-with-ent/ent/proto/entpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type handler struct {
	*ogent.OgentHandler
	client *ent.Client
}

func (h *handler) Health(ctx context.Context) (ogent.HealthNoContent, error) {
	return ogent.HealthNoContent{}, nil
}

func main() {
	// Create ent client.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	// Run the migrations.
	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Fatal(err)
	}

	// Start listening.
	srv, err := ogent.NewServer(&handler{ogent.NewOgentHandler(client), client})
	if err != nil {
		log.Fatal(err)
	}

	go pbserver(client)

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}

func pbserver(client *ent.Client) {
	svc := entpb.NewUserService(client)
	server := grpc.NewServer()
	entpb.RegisterUserServiceServer(server, svc)
	reflection.Register(server)
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
