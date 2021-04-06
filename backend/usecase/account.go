package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/domain/service"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Account interface {
	SignUp(authAccount AuthAccount) error
	SignIn(authAccount *AuthAccount) (string, error)
	SignOut(authAccount *AuthAccount) (*entity.Account, error)
	Verify(authAccount *AuthAccount) (*entity.Account, error)
	UpdatePassword(*entity.Account) (*AuthAccount, error)
	UpdateProfile(*entity.Account) (*entity.Account, error)
}

type account struct {
	repository repository.AccountRepository
	service    service.AccountService
	auth       service.AuthService
}


func (a account) Verify(authAccount *AuthAccount) (*entity.Account, error) {
	accountId, err := a.auth.Verification(authAccount.Token)
	if err != nil {
		return nil, err
	}
	ac, err := a.repository.FindById(accountId)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

func (a account) SignUp(e AuthAccount) error {
	ac, err := a.repository.CountById(e.ID)
	if ac > 0 {
		return errors.New("already exist")
	}

	password := []byte(e.Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, 10)
	e.Password = string(hashedPassword)

	_, err = a.repository.Create(&entity.Account{ID: e.ID, Password: e.Password})
	if err != nil {
		return err
	}
	return nil
}

func (a account) SignIn(e *AuthAccount) (string, error) {
	token, err := a.auth.Authorize(&entity.Account{ID: e.ID, Password: e.Password})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a account) SignOut(e *AuthAccount) (*entity.Account, error) {
	panic("implement me")
}

func (a account) UpdatePassword(e *entity.Account) (*AuthAccount, error) {
	panic("implement me")
}

func (a account) UpdateProfile(e *entity.Account) (*entity.Account, error) {
	panic("implement me")
}

func NewAccount(repository repository.AccountRepository, service service.AccountService, auth service.AuthService) Account {
	return &account{repository: repository, service: service, auth: auth}
}
