package service

import (
	"errors"
	"strings"
	"time"

	"task-manager/cmd/api/internals/model"
	"task-manager/cmd/api/internals/repository"
)

var (
	ErrEmptyTitle    = errors.New("title is required")
	ErrInvalidStatus = errors.New("status must be in_progress or complete")
	ErrInvalidDate   = errors.New("due date must be in YYYY-MM-DD format")
)
const dateLayout = "2006-01-02"

type CreateTaskInput struct {
	Title       string
	Description string
	Status      string // optional, defaults to in_progress
	DueDate     string // optional, YYYY-MM-DD
}

// UpdateTaskInput uses pointers so omitted fields are left unchanged.
type UpdateTaskInput struct {
	Title       *string
	Description *string
	Status      *string
	DueDate     *string // pointer to "" clears the due date
}

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func validStatus(s string) bool {
	return s == model.StatusInProgress || s == model.StatusComplete
}

func validDate(s string) bool {
	if s == "" {
		return true
	}
	_, err := time.Parse(dateLayout, s)
	return err == nil
}

func (s *TaskService) Create(in CreateTaskInput) (*model.Task, error) {
	title := strings.TrimSpace(in.Title)
	if title == "" {
		return nil, ErrEmptyTitle
	}
	status := in.Status
	if status == "" {
		status = model.StatusInProgress
	}
	if !validStatus(status) {
		return nil, ErrInvalidStatus
	}
	if !validDate(in.DueDate) {
		return nil, ErrInvalidDate
	}
	now := time.Now().UTC()
	return s.repo.Create(&model.Task{
		Title:       title,
		Description: strings.TrimSpace(in.Description),
		Status:      status,
		DueDate:     in.DueDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
}

func (s *TaskService) Get(id int64) (*model.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) List(status string) ([]model.Task, error) {
	if status != "" && !validStatus(status) {
		return nil, ErrInvalidStatus
	}
	return s.repo.List(status)
}

func (s *TaskService) Update(id int64, in UpdateTaskInput) (*model.Task, error) {
	t, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if in.Title != nil {
		title := strings.TrimSpace(*in.Title)
		if title == "" {
			return nil, ErrEmptyTitle
		}
		t.Title = title
	}
	if in.Description != nil {
		t.Description = strings.TrimSpace(*in.Description)
	}
	if in.Status != nil {
		if !validStatus(*in.Status) {
			return nil, ErrInvalidStatus
		}
		t.Status = *in.Status
	}
	if in.DueDate != nil {
		if !validDate(*in.DueDate) {
			return nil, ErrInvalidDate
		}
		t.DueDate = *in.DueDate
	}
	t.UpdatedAt = time.Now().UTC()
	if err := s.repo.Update(t); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *TaskService) Complete(id int64) (*model.Task, error) {
	status := model.StatusComplete
	return s.Update(id, UpdateTaskInput{Status: &status})
}

func (s *TaskService) Delete(id int64) error {
	return s.repo.Delete(id)
}
