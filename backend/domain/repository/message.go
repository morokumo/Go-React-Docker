package repository

import "backend/domain/entity"

type MessageRepository interface {
	CountById(id string) (int64, error)
	FindAll() (*[]entity.Message, error)
	FindById(id string) (*entity.Message, error)
	FindByRoom(room *entity.Room) (*[]entity.Message, error)
	Create(*entity.Message) (*entity.Message, error)
	Delete(*entity.Message) error
	Update(*entity.Message) (*entity.Message, error)
}
