package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/domain/service"
	"errors"
	"log"
	"time"
)

type Message interface {
	Verify(sendContent *SendContent) error
	FindMyRooms(account *entity.Account) (*[]responseRoom, error)
	FindPublicRooms(account *entity.Account) (*[]responseRoom, error)
	CreateRoom(room *entity.Room) error
	DeleteRoom(sendContent *SendContent) error
	JoinRoom(sendContent *SendContent) error
	LeaveRoom(sendContent *SendContent) error
	UpdateRoom(sendContent *SendContent) error
	SendMessage(sendContent *SendContent) (*responseMessage,error)
	GetMessageByRoom(sendContent *SendContent) (*[]responseMessage, error)
	GetMessageByAccount(account *entity.Account) error
}
type message struct {
	messageRepository repository.MessageRepository
	roomRepository    repository.RoomRepository
	messageService    service.MessageService
}

func (m message) FindPublicRooms(account *entity.Account) (*[]responseRoom, error) {
	rooms, err := m.roomRepository.FindPublic(account)
	if err != nil {

	}
	return convertResponseRoom(rooms), err
}

type responseRoom struct {
	ID      string
	Name    string
	Info    string
	Private bool
	CreateTime  time.Time
}
type responseMessage struct {
	AccountID string
	Text      string
	SendTime  time.Time
}

type SendContent struct {
	RoomId      string `json:"room_id"`
	MessageText string `json:"message"`
	Account     *entity.Account
}

func (m message) Verify(sendContent *SendContent) error {

	cnt, err := m.roomRepository.CountByAccountAndRoom(&entity.Account{ID: sendContent.Account.ID}, &entity.Room{ID: sendContent.RoomId})
	if err != nil {
		return err
	}
	if cnt == 0 {
		return errors.New("This room ID is not associated with an account.")
	}
	return nil
}

func (m message) DeleteRoom(sendContent *SendContent) error {
	panic("implement me")
}

func (m message) JoinRoom(sendContent *SendContent) error {
	return m.roomRepository.AddAccount(&entity.Room{ID: sendContent.RoomId},sendContent.Account)
}

func (m message) LeaveRoom(sendContent *SendContent) error {
	panic("implement me")
}

func (m message) UpdateRoom(sendContent *SendContent) error {
	panic("implement me")
}

func (m message) GetMessageByRoom(sendContent *SendContent) (*[]responseMessage, error) {
	messages, err := m.messageRepository.FindByRoom(&entity.Room{ID: sendContent.RoomId})
	var res []responseMessage
	for _, msg := range *messages {
		res = append(res, responseMessage{AccountID: msg.AccountID, Text: msg.Text, SendTime: msg.CreatedAt})
	}
	return &res, err
}

func (m message) GetMessageByAccount(account *entity.Account) error {
	panic("implement me")
}

func (m message) FindMyRooms(account *entity.Account) (*[]responseRoom, error) {
	log.Println("AAAAA",account)
	rooms, err := m.roomRepository.FindByAccount(account)
	if err != nil {

	}
	return convertResponseRoom(rooms), err
}

func (m message) CreateRoom(room *entity.Room) error {
	m.roomRepository.Create(room)
	return nil
}

func (m message) SendMessage(sendContent *SendContent) (*responseMessage,error) {
	room, err := m.roomRepository.FindById(sendContent.RoomId)
	_, err = m.roomRepository.CountById(sendContent.RoomId)
	if !(room.ID == sendContent.RoomId && err == nil) {
		return nil,errors.New("room not found")
	}
	msg, err := m.messageRepository.Create(&entity.Message{Text: sendContent.MessageText, AccountID: sendContent.Account.ID})
	err = m.roomRepository.AddMessage(room, msg)
	if err != nil {
		return nil,err
	}
	return &responseMessage{AccountID: sendContent.Account.ID,Text: msg.Text,SendTime: msg.CreatedAt},nil
}

func convertResponseRoom(rooms *[]entity.Room) *[]responseRoom {
	var res []responseRoom
	for _, room := range *rooms {
		res = append(res, responseRoom{ID: room.ID, Name: room.Name, Info: room.Info, Private: room.Private,CreateTime: room.CreatedAt})
	}
	return &res
}

func NewMessage(messageRepository repository.MessageRepository, roomRepository repository.RoomRepository, service service.MessageService) Message {
	return &message{messageRepository: messageRepository, roomRepository: roomRepository, messageService: service}
}
