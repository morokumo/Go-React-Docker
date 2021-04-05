package entity

import (
	"gorm.io/gorm"
	"time"
)

type Room struct {
	ID      string `gorm:"primary_key"`
	Name    string
	Info    string
	Private bool

	Accounts []Account `gorm:"many2many:room_accounts"`
	Messages []Message `gorm:"many2many:room_messages"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
