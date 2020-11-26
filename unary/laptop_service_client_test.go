package unary_service_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/gRPC/pb"
	"github.com/gRPC/sample"
	"github.com/gRPC/store"
	unary_service "github.com/gRPC/unary"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopServer, laptopServerAddress := startLaptopGrpcServer(t)
	client := newLaptopGrpcClient(t, laptopServerAddress)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	res, err := client.CreateLaptop(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.GetId())

	savedLaptop := laptopServer.Store.Find(res.GetId())
	require.NotNil(t, savedLaptop)
	require.Equal(t, laptop.Id, savedLaptop.Id)

	requireSameLaptop(t, laptop, savedLaptop)
}

func startLaptopGrpcServer(t *testing.T) (*unary_service.LaptopServiceServer, string) {
	imMemmory := store.NewInMemoryLaptopStore()
	laptopServer := unary_service.NewServer(imMemmory)

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":80")
	require.NoError(t, err)

	go grpcServer.Serve(listener)
	return laptopServer, listener.Addr().String()
}

func newLaptopGrpcClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	jsonLaptop1, err := protojson.Marshal(laptop1)
	require.NoError(t, err)

	jsonLaptop2, err := protojson.Marshal(laptop2)
	require.NoError(t, err)

	require.Equal(t, jsonLaptop1, jsonLaptop2)
}
