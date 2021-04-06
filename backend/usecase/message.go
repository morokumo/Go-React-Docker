package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/domain/service"
	"errors"
)

type Message interface {
	// アカウントとルームが紐づけられているかの検証
	Verify(requestMessage *RequestMessage) error
	FindMyRooms(account *entity.Account) (*[]responseRoom, error)
	FindPublicRooms(account *entity.Account) (*[]responseRoom, error)
	FindRoomAccounts(requestMessage *RequestMessage) (*[]responseAccount, error)

	CreateRoom(room *entity.Room) error
	DeleteRoom(requestMessage *RequestMessage) error
	JoinRoom(requestMessage *RequestMessage) error
	LeaveRoom(requestMessage *RequestMessage) error
	UpdateRoom(requestMessage *RequestMessage) error
	SendMessage(requestMessage *RequestMessage) (*responseMessage, error)
	GetMessageByRoom(requestMessage *RequestMessage) (*[]responseMessage, error)
	GetMessageByAccount(account *entity.Account) error
}

type message struct {
	messageRepository repository.MessageRepository
	roomRepository    repository.RoomRepository
	messageService    service.MessageService
}

func (m message) Verify(sendContent *RequestMessage) error {

	cnt, err := m.roomRepository.CountByAccountAndRoom(&entity.Account{ID: sendContent.Account.ID}, &entity.Room{ID: sendContent.RoomId})
	if err != nil {
		return err
	}
	if cnt == 0 {
		return errors.New("This room ID is not associated with an account.")
	}
	return nil
}
func (m message) FindMyRooms(account *entity.Account) (*[]responseRoom, error) {
	rooms, err := m.roomRepository.FindByAccount(account)
	if err != nil {
		return nil, err
	}
	return convertResponseRoom(rooms), err
}
func (m message) FindPublicRooms(account *entity.Account) (*[]responseRoom, error) {
	rooms, err := m.roomRepository.FindPublic(account)
	if err != nil {
		return nil, err
	}
	return convertResponseRoom(rooms), err
}

func (m message) FindRoomAccounts(requestMessage *RequestMessage) (*[]responseAccount, error) {
	accounts, err := m.roomRepository.FindRoomAccounts(requestMessage.RoomId)
	if err != nil {
		return nil, err
	}
	return convertResponseAccount(accounts), err
}

func (m message) DeleteRoom(sendContent *RequestMessage) error {
	panic("implement me")
}

func (m message) JoinRoom(sendContent *RequestMessage) error {
	return m.roomRepository.AddAccount(&entity.Room{ID: sendContent.RoomId}, sendContent.Account)
}

func (m message) LeaveRoom(sendContent *RequestMessage) error {
	panic("implement me")
}

func (m message) UpdateRoom(sendContent *RequestMessage) error {
	panic("implement me")
}

func (m message) GetMessageByRoom(sendContent *RequestMessage) (*[]responseMessage, error) {
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



func (m message) CreateRoom(room *entity.Room) error {
	_, err := m.roomRepository.Create(room)
	if err != nil {
		return err
	}
	return nil
}

func (m message) SendMessage(sendContent *RequestMessage) (*responseMessage, error) {
	room, err := m.roomRepository.FindById(sendContent.RoomId)
	if err != nil {
		return nil, err
	}

	msg, err := m.messageRepository.Create(&entity.Message{Text: sendContent.MessageText, AccountID: sendContent.Account.ID})
	if err != nil {
		return nil, err
	}

	err = m.roomRepository.AddMessage(room, msg)

	if err != nil {
		return nil, err
	}

	return &responseMessage{AccountID: sendContent.Account.ID, Text: msg.Text, SendTime: msg.CreatedAt}, nil
}


func NewMessage(messageRepository repository.MessageRepository, roomRepository repository.RoomRepository, service service.MessageService) Message {
	return &message{messageRepository: messageRepository, roomRepository: roomRepository, messageService: service}
}
