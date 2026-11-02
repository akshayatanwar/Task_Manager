package repository

import (
	"errors"

	"task-manager/cmd/api/internal/database"
	"task-manager/cmd/api/internal/models"
)

var ErrNotFound = errors.New("task not found")

type TaskRepository struct {
	store *database.Store
}

func NewTaskRepository(store *databse.Store) *TaskRepository {
	return &TaskRepository{store: store}
}

func (r *TaskRepository) Create(task models.Task) (models.Task, error) {

}

func (r *TaskRepository) GetAll() ([]models.Task, error) {

}

func (r *TaskRepository) GetByID(id int) (models.Task, error) {

}

func (r *TaskRepository) Update(id int, task models.Task) (models.Task, error) {

}

func (r *TaskRepository) Delete(id int) error {

}
