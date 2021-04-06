package dao

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db gorm.DB
}


func (r RoomRepository) FindPublic(account *entity.Account) (*[]entity.Room, error) {
	var rooms []entity.Room
	s := r.db.Table("room_accounts").Where("account_id = ?", account.ID).Select("room_id")
	r.db.Table("rooms").Where("id not in (?)", s).Where("private = ?", false).Find(&rooms)
	return &rooms, nil
}

func (r RoomRepository) FindRoomAccounts(id string) (*[]entity.Account, error) {
	var accounts []entity.Account
	s := r.db.Table("room_accounts").Where("room_id = ?", id).Select("account_id")
	r.db.Debug().Table("accounts").Where("id in (?)", s).Find(&accounts)

	//r.db.Preload("accounts").Where("id = ?", id).Find(&accounts)
	return &accounts, nil
}

func (r RoomRepository) CountByAccountAndRoom(account *entity.Account, room *entity.Room) (int64, error) {
	r.db.Preload("Rooms", "id = ?", room.ID).Find(account)
	return int64(len(account.Rooms)), nil
}

func (r RoomRepository) CountById(id string) (int64, error) {
	var cnt int64
	r.db.Model(&entity.Room{}).Where("id = ?", id).Count(&cnt)
	return cnt, nil
}

func (r RoomRepository) AddMessage(room *entity.Room, message *entity.Message) error {
	err := r.db.Model(&room).Association("Messages").Append(message)
	return err
}

func (r RoomRepository) AddAccount(room *entity.Room, account *entity.Account) error {
	err := r.db.Model(room).Association("Accounts").Append(account)
	return err
}

func (r RoomRepository) FindAll() (*[]entity.Room, error) {
	panic("implement me")
}

func (r RoomRepository) FindById(id string) (*entity.Room, error) {
	var res *entity.Room
	r.db.Where("id = ?", id).Find(&res)
	return res, nil
}

func (r RoomRepository) FindByAccount(account *entity.Account) (*[]entity.Room, error) {
	r.db.Preload("Rooms").Where("id = ?", account.ID).Find(&account)
	return &account.Rooms, nil
}

func (r RoomRepository) FindByMessage(message *entity.Message) (*entity.Room, error) {
	panic("implement me")
}

func (r RoomRepository) Create(room *entity.Room) (*entity.Room, error) {
	r.db.Create(room)
	return room, nil
}

func (r RoomRepository) Delete(room *entity.Room) error {
	panic("implement me")
}

func (r RoomRepository) Update(room *entity.Room) (*entity.Room, error) {
	panic("implement me")
}

func NewRoomRepository(db *gorm.DB) repository.RoomRepository {
	newRepository := &RoomRepository{
		db: *db,
	}
	return newRepository
}
