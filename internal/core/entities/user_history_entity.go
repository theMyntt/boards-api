package entities

import (
	"boards-api/internal/core/entities/enums"
	"time"

	"github.com/samborkent/uuidv7"
)

type UserHistoryEntity struct {
	ID          uuidv7.UUID      `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Owner       string           `json:"owner"`
	Responsible []string         `json:"responsible"`
	Status      enums.StatusEnum `json:"status"`
	SubTasks    []TaskEntity     `json:"sub_tasks"`
	CreatedAt   time.Time        `json:"created_at"`
}
