package registry

import (
	"backend/domain/service"
)

type Service interface {
	NewAccountService() service.AccountService
	NewMessageService() service.MessageService
	NewAuthService() service.AuthService
}

type serviceImpl struct {
	repository     Repository
	accountService service.AccountService
	messageService service.MessageService
	authService    service.AuthService
}

func (s serviceImpl) NewAccountService() service.AccountService {
	s.accountService = service.NewAccountService(s.repository.NewAccountRepository())
	return s.accountService
}

func (s serviceImpl) NewMessageService() service.MessageService {
	s.messageService = service.NewMessageService(s.repository.NewMessageRepository())
	return s.messageService
}
func (s serviceImpl) NewAuthService() service.AuthService {
	s.authService = service.NewAuthService(s.repository.NewAccountRepository())
	return s.authService
}

func NewService(repository Repository) Service {
	return &serviceImpl{repository: repository}
}
