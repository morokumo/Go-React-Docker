package repository

import "backend/domain/entity"

type RoomRepository interface {
	CountById(id string) (int64, error)
	CountByAccountAndRoom(account *entity.Account, room *entity.Room) (int64, error)
	FindAll() (*[]entity.Room, error)
	FindPublic(account *entity.Account) (*[]entity.Room, error)
	FindById(id string) (*entity.Room, error)
	FindRoomAccounts(id string) (*[]entity.Account, error)
	FindByAccount(account *entity.Account) (*[]entity.Room, error)
	FindByMessage(message *entity.Message) (*entity.Room, error)
	AddMessage(room *entity.Room, message *entity.Message) error
	AddAccount(room *entity.Room, account *entity.Account) error
	Create(room *entity.Room) (*entity.Room, error)
	Delete(room *entity.Room) error
	Update(room *entity.Room) (*entity.Room, error)
}
