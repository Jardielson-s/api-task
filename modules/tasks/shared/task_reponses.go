package shared

import "time"

type CreateTaskResponse struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"name"`
	Summary   string `gorm:"summary"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
