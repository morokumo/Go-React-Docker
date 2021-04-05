package dao

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db gorm.DB
}

func (a AccountRepository) FindAll() (*[]entity.Account, error) {
	panic("implement me")
}

func (a AccountRepository) FindById(id string) (*entity.Account, error) {
	var account entity.Account
	a.db.Where("id = ?", id).Find(&account)
	return &account, nil
}
func (a AccountRepository) CountById(id string) (int64, error) {
	var cnt int64
	a.db.Model(&entity.Account{}).Where("id = ?", id).Count(&cnt)
	return cnt, nil
}

func (a AccountRepository) FindByName(name string) (*[]entity.Account, error) {
	panic("implement me")
}

func (a AccountRepository) Create(account *entity.Account) (*entity.Account, error) {
	result := a.db.Create(account)
	return account, result.Error
}

func (a AccountRepository) Delete(account *entity.Account) error {
	panic("implement me")
}

func (a AccountRepository) Update(account *entity.Account) (*entity.Account, error) {
	panic("implement me")
}

func NewAccountRepository(db *gorm.DB) repository.AccountRepository {
	newRepository := &AccountRepository{
		db: *db,
	}
	return newRepository
}
