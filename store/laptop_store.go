package store

import (
	"errors"
	"sync"

	"github.com/gRPC/pb"
)

var ErrConflict error = errors.New("Laptop is already exists")

type LaptopStore interface {
	Save(*pb.Laptop) error
	Find(string) *pb.Laptop
	Search(*pb.FilterRequest, func(*pb.Laptop) error) error
}

type inMemoryLaptopStore struct {
	data map[string]*pb.Laptop
	mu   sync.RWMutex
}

func NewInMemoryLaptopStore() LaptopStore {
	return &inMemoryLaptopStore{
		data: map[string]*pb.Laptop{},
		mu:   sync.RWMutex{},
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

func (store *inMemoryLaptopStore) Find(id string) *pb.Laptop {
	store.mu.Lock()
	defer store.mu.Unlock()

	return store.data[id]
}

func (store *inMemoryLaptopStore) Search(filter *pb.FilterRequest, callback func(*pb.Laptop) error) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	for _, laptop := range store.data {
		if laptop.GetBrand() == filter.GetBrand() {
			err := callback(laptop)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
