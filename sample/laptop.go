package sample

import (
	"github.com/gRPC/pb"
	"github.com/google/uuid"
)

func NewLaptop() *pb.Laptop {
	return &pb.Laptop{Id: uuid.New().String()}
}
