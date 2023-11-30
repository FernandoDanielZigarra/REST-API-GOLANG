package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	ID          string `gorm:"primaryKey"`
	Title       string `gorm:"not null;type:varchar(100);unique_index;" json:"title"`
	Description string `json:"description"`
	Done        bool   `gorm:"default:false" json:"done"`
	UserID      string   `json:"user_id"`
}

func (task *Task) BeforeCreate(tx *gorm.DB) (err error) {
	task.ID = uuid.New().String()
	return
}
