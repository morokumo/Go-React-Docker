package service

import "backend/domain/repository"

type AccountService interface {
	SignIn()
	SignOut()
	UpdatePassword()
	UpdateProfile()
}
type accountService struct {
	repository repository.AccountRepository
}

func (a accountService) SignIn() {
	panic("implement me")
}

func (a accountService) SignOut() {
	panic("implement me")
}

func (a accountService) UpdatePassword() {
	panic("implement me")
}

func (a accountService) UpdateProfile() {
	panic("implement me")
}

func NewAccountService(repository repository.AccountRepository) AccountService {
	return &accountService{repository: repository}
}