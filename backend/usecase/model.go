package usecase

import (
	"backend/domain/entity"
	"time"
)

type AuthAccount struct {
	ID       string `json:"account_id"`
	Password string `json:"password"`
	Token    string
}

type RequestMessage struct {
	RoomId      string `json:"room_id"`
	MessageText string `json:"message"`
	Account     *entity.Account
}

type RequestRoom struct {
	Name     string `json:"room_name"`
	Info     string `json:"info"`
	Private  string `json:"private"`
	Password string `json:"password"`
}

type responseRoom struct {
	ID         string
	Name       string
	Info       string
	Private    bool
	CreateTime time.Time
}

type responseMessage struct {
	AccountID string
	Text      string
	SendTime  time.Time
}
