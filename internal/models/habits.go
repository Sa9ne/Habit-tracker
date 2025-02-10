package models

import (
	"time"
)

type Habits struct {
	ID                 uint `gorm:"primaryKey"`
	HabitName          string
	CompletedTasks     int
	MoreInfo           string
	DifficultiLevel    int
	IsFireStreakActive bool
	LastCompletedDate  time.Time `gorm:"type:date"` // Добавляем поле для даты последнего выполнения
}
