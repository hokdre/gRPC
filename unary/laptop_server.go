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

type LaptopServiceServer struct {
	Store store.LaptopStore
	pb.UnimplementedLaptopServiceServer
}

func NewServer(store store.LaptopStore) *LaptopServiceServer {
	return &LaptopServiceServer{
		Store: store,
	}
}

func (service *LaptopServiceServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive request create laptop with id : %s \n", laptop.Id)

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

	if err := service.Store.Save(laptop); err != nil {
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

func (service *LaptopServiceServer) SearchLaptop(
	req *pb.FilterRequest,
	stream pb.LaptopService_SearchLaptopServer,
) error {
	log.Printf("receive request search laptop with brand : %s \n", req.GetBrand())

	err := service.Store.Search(req, func(laptop *pb.Laptop) error {
		if err := stream.Send(laptop); err != nil {
			return err
		}

		log.Printf("sent laptop with id : %s \n", laptop.GetId())
		return nil
	})

	if err != nil {
		return status.Errorf(codes.Internal, "unxecpted error : %s", err)
	}

	return nil
}
