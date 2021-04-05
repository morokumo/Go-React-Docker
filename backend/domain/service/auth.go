package service

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"errors"
	"github.com/form3tech-oss/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService interface {
	Authenticate(account *entity.Account) (*entity.Account, error)
	Authorize(account *entity.Account) (string, error)
	Verification(token string) (string, error)
}

type authService struct {
	repository repository.AccountRepository
	signature  []byte
}

func (a authService) Authenticate(account *entity.Account) (*entity.Account, error) {
	ac, err := a.repository.FindById(account.ID)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(ac.Password), []byte(account.Password))
	if err != nil {
		return nil, err
	}
	return ac, err
}

func (a authService) Authorize(account *entity.Account) (string, error) {
	ac, err := a.Authenticate(account)
	if ac == nil || err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["subject"] = account.ID
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()
	tokenStr, _ := token.SignedString(a.signature)
	return tokenStr, nil
}

func (a authService) Verification(tokenStr string) (string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return a.signature, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	if _, ok := claims["subject"]; !ok {
		return "", errors.New("key not found.")
	}
	accountID := claims["subject"].(string)
	return accountID, nil
}

func NewAuthService(accountRepository repository.AccountRepository) AuthService {
	signature := []byte("こんにちは!")
	return &authService{accountRepository, signature}
}
