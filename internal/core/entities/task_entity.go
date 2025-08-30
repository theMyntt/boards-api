package entities

import (
	"boards-api/internal/core/entities/enums"
	"time"

	"github.com/samborkent/uuidv7"
)

type TaskEntity struct {
	ID          uuidv7.UUID      `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Owner       string           `json:"owner"`
	Responsible string           `json:"responsible"`
	Status      enums.StatusEnum `json:"status"`
	CreatedAt   time.Time        `json:"created_at"`
}
