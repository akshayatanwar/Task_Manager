package repository

import (
	"errors"

	"task-manager/cmd/api/internal/database"
	"task-manager/cmd/api/internal/models"
)

var ErrNotFound = errors.New("task not found")

type Task
