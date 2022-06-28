package main

import (
	"context"
	"fmt"
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
	ms "github.com/ryota0624/multi-server"
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

	restApiPort, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServicesPortNumber := 5000
	grpcServicesPort, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcServicesPortNumber))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcuiPort, err := net.Listen("tcp", fmt.Sprintf(":%d", 5005))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	servers := ms.NewServers().
		Resister(
			ms.NewHttpServer(serveGrpcui(grpcServicesPortNumber), grpcuiPort),
		).
		Resister(
			ms.NewHttpServer(&http.Server{
				Handler: srv,
			}, restApiPort),
		).
		Resister(
			ms.NewGrpcServer(serveGrpcServices(client), grpcServicesPort),
		).
		EnableShutdownOnTerminateSignal().
		ShutdownTimout(time.Second * 3)

	log.Println("start servers")

	err = servers.Start(context.Background())
	if err != nil {
		log.Printf("occurred servers start err: %v\n", err)
	}
	log.Println("shutdown start")
	servers.WaitShutdown()
}

func serveGrpcServices(client *ent.Client) *grpc.Server {
	svc := entpb.NewUserService(client)
	server := grpc.NewServer()
	entpb.RegisterUserServiceServer(server, svc)
	reflection.Register(server)
	return server
}

func serveGrpcui(grpcServicesPort int) *http.Server {
	ctx := context.Background()
	dialTime := 10 * time.Second
	dialCtx, cancel := context.WithTimeout(ctx, dialTime)
	defer cancel()

	grpcServicesAddr := fmt.Sprintf("127.0.0.1:%d", grpcServicesPort)
	time.Sleep(time.Second * 3) // grpc serverの起動を待つ
	var creds credentials.TransportCredentials
	clientConn, err := grpcurl.BlockingDial(dialCtx, "tcp", grpcServicesAddr, creds)
	if err != nil {
		log.Fatalf("Failed to dial target host : %+v", err)
	}

	handler, err := standalone.HandlerViaReflection(ctx, clientConn, grpcServicesAddr)
	if err != nil {
		log.Fatalf("failed to HandlerViaReflection: %s", err)
	}

	return &http.Server{
		Handler: handler,
	}
}
