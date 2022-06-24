package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/fullstorydev/grpcui/standalone"
	"github.com/fullstorydev/grpcurl"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ryota0624/go-dev-with-ent/ent"
	"github.com/ryota0624/go-dev-with-ent/ent/ogent"
	"github.com/ryota0624/go-dev-with-ent/ent/proto/entpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	go func() {
		go serveGrpcServices(client)
		go serveGrpcui()
	}()
	log.Default().Println("start REST API server at 8080")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}

func serveGrpcServices(client *ent.Client) {
	svc := entpb.NewUserService(client)
	server := grpc.NewServer()
	entpb.RegisterUserServiceServer(server, svc)
	reflection.Register(server)
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}
	log.Default().Println("start grpc server at 5000")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}

func serveGrpcui() {
	ctx := context.Background()
	dialTime := 10 * time.Second
	dialCtx, cancel := context.WithTimeout(ctx, dialTime)
	defer cancel()

	time.Sleep(time.Second * 1) // grpc serverの起動を待つ
	var creds credentials.TransportCredentials
	clientConn, err := grpcurl.BlockingDial(dialCtx, "tcp", "127.0.0.1:5000", creds)
	if err != nil {
		log.Fatalf("Failed to dial target host : %+v", err)
	}

	handler, err := standalone.HandlerViaReflection(ctx, clientConn, "127.0.0.1:5000")
	if err != nil {
		log.Fatalf("failed to HandlerViaReflection: %s", err)
	}

	log.Default().Println("start grpcui server at 5005")
	err = http.ListenAndServe(":5005", handler)
	if err != nil {
		log.Fatalf("grpcui server ended: %s", err)
	}
}
