package entity

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID       string `gorm:"primary_key"`
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Profile

	Rooms []Room `gorm:"many2many:room_accounts"`
}

type Profile struct {
	Name string
	Age  int64
}
