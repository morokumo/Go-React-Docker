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

func convertResponseRoom(rooms *[]entity.Room) *[]responseRoom {
	var res []responseRoom
	for _, room := range *rooms {
		res = append(res, responseRoom{ID: room.ID, Name: room.Name, Info: room.Info, Private: room.Private, CreateTime: room.CreatedAt})
	}
	return &res
}

type responseMessage struct {
	AccountID string
	Text      string
	SendTime  time.Time
}
type responseAccount struct {
	ID string
}
func convertResponseAccount(accounts *[]entity.Account) *[]responseAccount {
	var res []responseAccount
	for _, account := range *accounts {
		res = append(res, responseAccount{ID: account.ID})
	}
	return &res
}