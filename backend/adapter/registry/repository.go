package registry

import (
	"backend/domain/repository"
	"backend/infra/dao"
	"gorm.io/gorm"
)

type Repository interface {
	NewAccountRepository() repository.AccountRepository
	NewMessageRepository() repository.MessageRepository
	NewRoomRepository() repository.RoomRepository
}

type repositoryImpl struct {
	database          *gorm.DB
	accountRepository repository.AccountRepository
	messageRepository repository.MessageRepository
	roomRepository    repository.RoomRepository
}

func (r repositoryImpl) NewRoomRepository() repository.RoomRepository {
	r.roomRepository = dao.NewRoomRepository(r.database)
	return r.roomRepository
}

func (r repositoryImpl) NewAccountRepository() repository.AccountRepository {
	r.accountRepository = dao.NewAccountRepository(r.database)
	return r.accountRepository
}

func (r repositoryImpl) NewMessageRepository() repository.MessageRepository {
	r.messageRepository = dao.NewMessageRepository(r.database)
	return r.messageRepository
}

func NewRepository() Repository {
	db, err := dao.ConnectDB()
	if err != nil {
		panic("cannot connect database.")
	}
	return &repositoryImpl{
		database: db,
	}
}
