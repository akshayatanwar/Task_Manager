package repository

import (
	"errors"
	"sort"

	"task-manager/cmd/api/internal/database"
	"task-manager/cmd/api/internal/models"
)

var ErrNotFound = errors.New("task not found")

type TaskRepository struct {
	store *database.Store
}

func NewTaskRepository(store *database.Store) *TaskRepository {
	return &TaskRepository{store: store}
}

func (r *TaskRepository) Create(task models.Task) (models.Task, error) {
	r.store.Mu.Lock()
	defer r.store.Mu.Unlock()

	task.ID = r.store.NextID
	r.store.NextID++
	r.store.Tasks[task.ID] = task

	if err := r.store.Save(); err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	r.store.Mu.RLock()
	defer r.store.Mu.RUnlock()

	tasks := make([]models.Task, 0, len(r.store.Tasks))
	for _, t := range r.store.Tasks {
		tasks = append(tasks, t)
	}
	sort.Slice(tasks, func(i, j int) bool { return tasks[i].ID < tasks[j].ID })
	return tasks, nil
}

func (r *TaskRepository) GetByID(id int) (models.Task, error) {
	r.store.Mu.RLock()
	defer r.store.Mu.RUnlock()

	task, ok := r.store.Tasks[id]
	if !ok {
		return models.Task{}, ErrNotFound
	}
	return task, nil
}

func (r *TaskRepository) Update(id int, task models.Task) (models.Task, error) {
	r.store.Mu.Lock()
	defer r.store.Mu.Unlock()

	if _, ok := r.store.Tasks[id]; !ok {
		return models.Task{}, ErrNotFound
	}
	task.ID = id
	r.store.Tasks[id] = task

	if err := r.store.Save(); err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepository) Delete(id int) error {
	r.store.Mu.Lock()
	defer r.store.Mu.Unlock()

	if _, ok := r.store.Tasks[id]; !ok {
		return ErrNotFound
	}
	delete(r.store.Tasks, id)
	return r.store.Save()
}
