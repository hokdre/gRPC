package main

import (
	"context"
	"flag"
	"fmt"
	"io"
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
	grpcOperation := flag.String("grpcType", "", "grpc operation : unary, sStream, cStream, bStream")
	flag.Parse()

	log.Printf("dial server on port : %d \n", *port)
	serverAddress := fmt.Sprintf("0.0.0.0:%d", *port)
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed dial server : %s", err)
	}

	client := pb.NewLaptopServiceClient(conn)

	flag.Parse()
	switch *grpcOperation {
	case "unary":
		createLaptop(client)
	case "sStream":
		searchLaptop(client)
	default:
		log.Fatalln("Please, Provided Operation!")
	}
}

func createLaptop(client pb.LaptopServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	laptop := sample.NewLaptop()
	laptop.Id = "invalid-uuid"
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	res, err := client.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok && st.Code() == codes.AlreadyExists {
			log.Print("cannot create laptop already exists")
		} else {
			log.Fatalf("cannot create laptop : %s", err)
		}
		return
	}

	log.Printf("success create laptop with id : %s", res.GetId())
}

func searchLaptop(client pb.LaptopServiceClient) {
	err := seedLaptop(client)
	if err != nil {
		log.Fatalf("seed laptop error : %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req := &pb.FilterRequest{
		Brand: "Even Brand",
	}

	stream, err := client.SearchLaptop(ctx, req)
	for {
		laptop, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("unxexpected error : %s", err)
		}

		log.Println("Receive Laptop :")
		log.Printf("Laptop ID: %s", laptop.GetId())
		log.Printf("Laptop Brand : %s", laptop.GetBrand())
	}
}

func seedLaptop(client pb.LaptopServiceClient) error {
	evenNumBrand := "Even Brand"
	oddNumBrand := "Odd Brand"

	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		laptop := sample.NewLaptop()
		if isOdd(i) {
			laptop.Brand = oddNumBrand
		} else {
			laptop.Brand = evenNumBrand
		}
		req := &pb.CreateLaptopRequest{
			Laptop: laptop,
		}

		res, err := client.CreateLaptop(ctx, req)
		if err != nil {
			return err
		}

		log.Printf("success create laptop with id : %s", res.GetId())
	}
	return nil
}

func isOdd(number int) bool {
	return (number % 2) == 1
}
