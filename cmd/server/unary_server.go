package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/gRPC/pb"
	"github.com/gRPC/store"
	unary_service "github.com/gRPC/unary"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 8080, "laptop server port")
	flag.Parse()

	laptopServer := unary_service.NewServer(store.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	serverAddress := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("failed create listener server : %s", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed serve grpc server : %s", err)
	}
}
