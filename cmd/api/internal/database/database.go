package database

import (
	"sync"

	"task-manager/cmd/api/internal/models"
)

type Store struct {
	mu     sync.RWMutex
	tasks  map[int]models.Task
	nextID int
	file   string
}

func NewStore(file string) (*Store, error) {
	s := &Store{
		tasks:  make(map[int]models.Task),
		nextID: 1,
		file:   file,
	}

	return s, nil
}
