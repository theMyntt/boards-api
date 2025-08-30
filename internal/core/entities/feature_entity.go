package entities

import (
	"boards-api/internal/core/entities/enums"
	"time"

	"github.com/samborkent/uuidv7"
)

type FeatureEntity struct {
	ID          uuidv7.UUID         `json:"id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Owner       string              `json:"owner"`
	Status      enums.StatusEnum    `json:"status"`
	Responsible []string            `json:"responsible"`
	SubTasks    []UserHistoryEntity `json:"sub_tasks"`
	CreatedAt   time.Time           `json:"created_at"`
}
