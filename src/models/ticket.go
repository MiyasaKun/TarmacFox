package models

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	ChannelID string    `gorm:"not null"`
	UserID    string    `gorm:"not null"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}




