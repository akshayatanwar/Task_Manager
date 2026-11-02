package database

import (
	"encoding/json"
	"os"
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
	if err := s.load(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Store) Save() error {
	list := make([]models.Task, 0, len(s.Tasks))
	for _, t := range s.Tasks {
		list = append(list, t)
	}
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.file, data, 0644)
}

func (s *Store) load() error {
	data, err := os.ReadFile(s.file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(data) == 0 {
		return nil
	}

	var list []models.Task
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}
	for _, t := range list {
		s.Tasks[t.ID] = t
		if t.ID >= s.NextID {
			s.NextID = t.ID + 1
		}
	}
	return nil
}
