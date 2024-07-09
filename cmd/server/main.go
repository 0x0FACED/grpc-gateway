package main

import (
	"context"
	"log"
	"net"
	"net/http"

	api "grpc-gateway/api/generated"
	"grpc-gateway/pkg/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterHostServiceServer(s, &service.HostService{})
	reflection.Register(s)

	go func() {
		log.Printf("Starting gRPC server on :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = api.RegisterHostServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	log.Printf("Starting HTTP server on :8080")
	if err := gwServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
