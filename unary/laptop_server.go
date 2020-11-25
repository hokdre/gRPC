package unary_service

import (
	"context"
	"errors"
	"log"

	"github.com/gRPC/pb"
	"github.com/gRPC/store"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type laptopServiceServer struct {
	store store.LaptopStore
	pb.UnimplementedLaptopServiceServer
}

func NewServer(store store.LaptopStore) pb.LaptopServiceServer {
	return &laptopServiceServer{
		store: store,
	}
}

func (service *laptopServiceServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	if laptop.Id != "" {
		if _, err := uuid.Parse(laptop.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "cannot parsed new id for laptop: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generated new id for laptop: %v", err)
		}
		laptop.Id = id.String()
	}

	if err := service.store.Save(laptop); err != nil {
		code := codes.Internal
		if errors.Is(err, store.ErrConflict) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop : %v", err)
	}

	log.Printf("saved laptop with id : %s \n", laptop.Id)
	response := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return response, nil
}
