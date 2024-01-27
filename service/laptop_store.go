package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/renatospaka/pc-book/pb"
)

var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	Save (laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data map[string]*pb.Laptop
}

// type DBLaptopStore struct {
	
// }

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data:  make(map[string]*pb.Laptop),
	}
}

func (db *InMemoryLaptopStore) Save (laptop *pb.Laptop) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if db.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	//deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %v", err)
	}

	db.data[other.Id] = other
	return nil
}

func (db *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	laptop := db.data[id]
	if laptop == nil {
		return nil, nil
	}

	// deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}
	return other, nil
}
