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
	rdb           *redis.Client
	tr            *repositories.TaskRepository
	durationCache time.Duration
}

func NewTaskService(
	rdb *redis.Client,
	tr *repositories.TaskRepository,
) *TaskService {
	return &TaskService{rdb, tr, time.Duration(time.Second * 10)}
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

	// TODO: melhorar isso
	// se tiver usuários, remove só do usuário pelo id dele
	// pois terá uma task nova
	// Limpa todos os bancos de dados do Redis
	err = ts.rdb.FlushAll(ctx).Err()
	if err != nil {
		fmt.Println("Erro ao limpar o cache Redis:", err)
	} else {
		fmt.Println("Cache Redis limpo com sucesso!")
	}

	// add cache by task id
	err = ts.rdb.Set(ctx, task.ID.String(), string(b), ts.durationCache).Err()
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

	// TODO: melhorar isso
	// se tiver usuários, remove só do usuário pelo id dele
	// pois terá uma task nova
	// Limpa todos os bancos de dados do Redis
	err = ts.rdb.FlushAll(ctx).Err()
	if err != nil {
		fmt.Println("Erro ao limpar o cache Redis:", err)
	} else {
		fmt.Println("Cache Redis limpo com sucesso!")
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
	err = ts.rdb.Set(ctx, task.ID.String(), string(b), ts.durationCache).Err()
	if err != nil {
		return nil, ErrCache
	}

	return task, nil

}

func (ts *TaskService) FindTasks(ctx context.Context, page, size int64, query string) ([]*models.Task, error) {
	// procura no cache
	var keyCache = fmt.Sprintf("tasks:%d:%d:%s", page, size, query)
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

	// procura no repository
	tasks, err := ts.tr.FindTasks(page, size, query)
	if err != nil {
		// manda pro logger
		return nil, ErrRepository
	}

	// grava no cache
	tasksBytes, err := json.Marshal(tasks)
	if err != nil {
		// manda pro logger
		return nil, ErrRepository
	}
	_, err = ts.rdb.Set(ctx, keyCache, string(tasksBytes), ts.durationCache).Result()
	if err != nil {
		// manda pro logger
		return nil, ErrCache
	}

	return tasks, nil
}
