package entity

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	AccountID string
	Text      string

	Rooms []Room `gorm:"many2many:room_messages"`
}
