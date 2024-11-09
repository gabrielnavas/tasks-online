package services

import (
	"api/models"
	"api/repositories"
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type TaskService struct {
	rdb *redis.Client
	tr  *repositories.TaskRepository
}

func NewTaskService(
	rdb *redis.Client,
	tr *repositories.TaskRepository,
) *TaskService {
	return &TaskService{rdb, tr}
}

func (ts *TaskService) CreateTask(ctx context.Context, description string) (*models.Task, error) {
	task, err := ts.tr.CreateTask(ctx, description)
	if err != nil {
		return nil, ErrRepository
	}

	b, err := json.Marshal(task)
	if err != nil {
		return nil, ErrCache
	}
	err = ts.rdb.Set(ctx, task.ID.String(), string(b), 0).Err()
	if err != nil {
		return nil, ErrCache
	}

	return task, nil
}

func (ts *TaskService) UpdateTask(ctx context.Context, taskId uuid.UUID, description string, done bool) error {
	return nil
}

func (ts *TaskService) DeleteTask(ctx context.Context, taskId uuid.UUID) error {
	return nil
}

func (ts *TaskService) FindTaskById(ctx context.Context, taskId uuid.UUID) (*models.Task, error) {
	taskJson, err := ts.rdb.Get(ctx, taskId.String()).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, ErrCache
	} else if len(taskJson) != 0 {
		// return cache task
		var task models.Task
		json.Unmarshal([]byte(taskJson), &task)
		return &task, nil
	}

	task, err := ts.tr.FindTaskById(ctx, taskId)
	if err != nil {
		return nil, ErrRepository
	}

	// add cache task
	b, err := json.Marshal(task)
	if err != nil {
		return nil, ErrCache
	}
	err = ts.rdb.Set(ctx, task.ID.String(), string(b), 0).Err()
	if err != nil {
		return nil, ErrCache
	}
	return task, nil

}

func (ts *TaskService) FindTasks(ctx context.Context, offset, size int64, query string) ([]*models.Task, error) {
	tasks, err := ts.tr.FindTasks(offset, size, query)
	if err != nil {
		// manda pro logger
		return nil, ErrRepository
	}
	return tasks, nil
}
