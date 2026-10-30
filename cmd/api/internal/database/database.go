package database

import (
	"sync"

	"task-manager/cmd/api/internal/models"
)

type Store struct {
	Mu     sync.RWMutex
	Tasks  map[int]models.Task
	NextID int
	file   string
}

func NewStore(file string) (*Store, error) {
	s := &Store{
		Tasks:  make(map[int]models.Task),
		NextID: 1,
		file:   file,
	}

	return s, nil
}
