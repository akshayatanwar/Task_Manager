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
