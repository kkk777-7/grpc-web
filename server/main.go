package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/kkk777-7/grpc-web/handler"
	pb "github.com/kkk777-7/grpc-web/messenger/api/v1"
)

const port = ":9090"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	pb.RegisterMessengerServer(srv, handler.NewMessengerHandler())
	reflection.Register(srv)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		srv.Serve(lis)
	}()

	// Graceful shutdown when Enter Ctrl+C.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	srv.GracefulStop()
}
