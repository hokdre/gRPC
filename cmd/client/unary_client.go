package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/gRPC/pb"
	"github.com/gRPC/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	port := flag.Int("port", 8080, "laptop server port")
	flag.Parse()
	log.Printf("dial server on port : %d \n", *port)

	serverAddress := fmt.Sprintf("0.0.0.0:%d", *port)
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed dial server : %s", err)
	}

	client := pb.NewLaptopServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req := &pb.CreateLaptopRequest{
		Laptop: sample.NewLaptop(),
	}
	res, err := client.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok && st.Code() == codes.AlreadyExists {
			log.Print("cannot create laptop already exists")
		} else {
			log.Fatalf("cannot create laptop : %s", err)
		}
	}

	log.Printf("success create laptop with id : %s", res.GetId())
}
