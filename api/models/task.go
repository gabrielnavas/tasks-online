package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Description string     `json:"description"`
	Done        bool       `json:"done"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

func (t *Task) Validate() error {
	if len(t.Description) < 2 {
		return errors.New("descrição deve ter pelo menos 2 caracteres")
	}
	if len(t.Description) > 100 {
		return errors.New("descrição deve ter no máximo 100 caracteres")
	}
	return nil
}
