package store

import (
	"errors"
	"sync"

	"github.com/gRPC/pb"
)

var ErrConflict error = errors.New("Laptop is already exists")

type LaptopStore interface {
	Save(*pb.Laptop) error
}

type inMemoryLaptopStore struct {
	data map[string]*pb.Laptop
	mu   sync.RWMutex
}

func NewInMemoryLaptopStore() LaptopStore {
	return &inMemoryLaptopStore{
		mu: sync.RWMutex{},
	}
}

func (store *inMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	isLaptopAlreadyExist := store.data[laptop.Id] != nil
	if isLaptopAlreadyExist {
		return ErrConflict
	}

	store.data[laptop.Id] = laptop
	return nil
}
