package shared

import "time"

type CreateUserResponse struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"username"`
	Email     string `gorm:"index:,unique"`
	Password  string `gorm:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
