package services

import (
	"api/models"
	"api/repositories"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

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
	err := ts.tr.DeleteTask(ctx, taskId)
	if err != nil {
		// log error
		return ErrRepository
	}

	_, err = ts.rdb.Del(ctx, taskId.String()).Result()
	if err != nil {
		// log error
		return ErrCache
	}
	return nil
}

func (ts *TaskService) FindTaskById(ctx context.Context, taskId uuid.UUID) (*models.Task, error) {
	taskJson, err := ts.rdb.Get(ctx, taskId.String()).Result()

	// any error
	if err != nil && !errors.Is(err, redis.Nil) {
		// log error
		return nil, ErrCache
	} else if len(taskJson) != 0 {
		// task found on cache
		var task models.Task
		json.Unmarshal([]byte(taskJson), &task)
		return &task, nil
	}

	// get from repository
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
	var keyCache = fmt.Sprintf("tasks:%d:%d:%s", offset, size, query)
	tasksJson, err := ts.rdb.Get(ctx, keyCache).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		// manda pro logger
		return nil, ErrCache
	}
	if len(tasksJson) != 0 {
		var tasks []*models.Task
		json.Unmarshal([]byte(tasksJson), &tasks)
		return tasks, nil
	}

	tasks, err := ts.tr.FindTasks(offset, size, query)
	if err != nil {
		// manda pro logger
		return nil, ErrRepository
	}

	tasksBytes, err := json.Marshal(tasks)
	_, err = ts.rdb.Set(ctx, keyCache, string(tasksBytes), time.Duration(time.Second*10)).Result()
	if err != nil {
		// manda pro logger
		return nil, ErrCache
	}
	return tasks, nil
}
