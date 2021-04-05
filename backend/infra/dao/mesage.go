package dao

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db gorm.DB
}

func (m MessageRepository) CountById(id string) (int64, error) {
	panic("implement me")
}

func (m MessageRepository) FindAll() (*[]entity.Message, error) {
	panic("implement me")
}

func (m MessageRepository) FindById(id string) (*entity.Message, error) {

	panic("implement me")
}

func (m MessageRepository) FindByRoom(room *entity.Room) (*[]entity.Message, error) {

	m.db.Preload("Messages").Where("id = ?",room.ID).Find(&room)
	return &room.Messages, nil
}

func (m MessageRepository) Create(message *entity.Message) (*entity.Message, error) {
	result := m.db.Create(message)
	return message, result.Error
}

func (m MessageRepository) Delete(message *entity.Message) error {
	panic("implement me")
}

func (m MessageRepository) Update(message *entity.Message) (*entity.Message, error) {
	panic("implement me")
}

func NewMessageRepository(db *gorm.DB) repository.MessageRepository {
	newRepository := &MessageRepository{
		db: *db,
	}
	return newRepository
}
