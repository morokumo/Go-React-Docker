package service

import (
	"backend/domain/entity"
	"backend/domain/repository"
)

type MessageService interface {
	Send(*entity.Message) error
	Get(*entity.Message) error
}

type messageService struct {
	repository repository.MessageRepository
}

func (m messageService) Send(message *entity.Message) error {
	panic("implement me")
}

func (m messageService) Get(message *entity.Message) error {
	panic("implement me")
}

func NewMessageService(repository repository.MessageRepository) MessageService {
	return &messageService{repository: repository}
}
