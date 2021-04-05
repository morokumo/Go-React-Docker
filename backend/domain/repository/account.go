package repository

import "backend/domain/entity"

type AccountRepository interface {
	FindAll() (*[]entity.Account, error)
	FindById(id string) (*entity.Account, error)
	CountById(id string) (int64, error)
	FindByName(name string) (*[]entity.Account, error)
	Create(*entity.Account) (*entity.Account, error)
	Delete(*entity.Account) error
	Update(*entity.Account) (*entity.Account, error)
}
