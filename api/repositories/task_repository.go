package repositories

import (
	"api/dtos"
	"api/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (tr *TaskRepository) CreateTask(ctx context.Context, description string) (*models.Task, error) {
	var t models.Task
	t.ID = uuid.New()
	t.Description = description
	t.Done = false
	t.CreatedAt = time.Now().UTC()

	statement, err := tr.db.PrepareContext(ctx, `
		INSERT INTO tasks (id, description, done, created_at)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	_, err = statement.ExecContext(ctx, t.ID, t.Description, t.Done, t.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (tr *TaskRepository) UpdateTask(taskId uuid.UUID, dto dtos.UpdateTaskDto) error {
	statement, err := tr.db.Prepare(`
	UPDATE tasks
	SET 
		done = ?,
		description = ?,
		updated_at = ?
	WHERE id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()
	now := time.Now().UTC()
	_, err = statement.Exec(dto.Done, dto.Description, now, taskId)
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) DeleteTask(ctx context.Context, taskId uuid.UUID) error {
	statement, err := tr.db.PrepareContext(ctx, `
		DELETE FROM tasks
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.ExecContext(ctx, taskId)
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) FindTaskById(ctx context.Context, taskId uuid.UUID) (*models.Task, error) {
	statement, err := tr.db.PrepareContext(ctx, `
		SELECT id, description, done
		FROM tasks
		WHERE id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	row := statement.QueryRowContext(ctx, taskId)

	var t models.Task
	err = row.Scan(&t.ID, &t.Description, &t.Done)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrResourceNotFound
		} else if errors.Is(err, sql.ErrConnDone) {
			return nil, ErrConnectionDone
		}
		return nil, err
	}

	return &t, nil
}

func (tr *TaskRepository) FindTasks(page, size int64, description string) ([]*models.Task, error) {
	page = (page - 1) * size
	tasks := []*models.Task{}
	statement, err := tr.db.Prepare(`
		SELECT id, description, done, created_at, updated_at
		FROM tasks
		WHERE
			(length(?) = 0 OR lower(description) LIKE lower(?))
		ORDER BY created_at DESC
		LIMIT ?
		OFFSET ?
	`)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	searchTerm := fmt.Sprintf("%%%s%%", description)
	rows, err := statement.Query(description, searchTerm, size, page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Task
		var updatedAt sql.NullTime

		if err := rows.Scan(&t.ID, &t.Description, &t.Done, &t.CreatedAt, &updatedAt); err != nil {
			return nil, err
		}

		if updatedAt.Valid {
			t.UpdatedAt = &updatedAt.Time
		} else {
			t.UpdatedAt = nil
		}

		tasks = append(tasks, &t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
