package handler

import (
	"backend/adapter/registry"
	"backend/usecase"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
)

type AccountHandler interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
	SignOut(ctx *gin.Context)
	Verify(ctx *gin.Context) error
	UpdatePassword(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type accountHandler struct {
	uc usecase.Account
}

func (a accountHandler) Verify(ctx *gin.Context) error {
	bodyCopy := new(bytes.Buffer)
	_, err := io.Copy(bodyCopy, ctx.Request.Body)
	if err != nil {
		BadRequest(ctx, err)
		return err
	}
	bodyData := bodyCopy.Bytes()
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
	account, err := convertAccount(ctx)

	if err != nil {
		BadRequest(ctx, err)
		return err
	}
	ac, err := a.uc.Verify(account)
	if ac ==nil || ac.ID == "" {
		UnAuthorized(ctx)
		return errors.New("account not found")
	}
	log.Println("POP", ac)
	if err != nil {
		UnAuthorized(ctx)
		return err
	}
	ctx.Set("account", ac)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
	return nil
}

func (a accountHandler) SignUp(ctx *gin.Context) {
	account, err := convertAccount(ctx)
	if err != nil {
		BadRequest(ctx, err)
		log.Println(err)
		return
	}
	err = a.uc.SignUp(*account)
	if err != nil {
		BadRequest(ctx, err)
		log.Println(err)
		return
	}
	token, err := a.uc.SignIn(account)
	if err != nil {
		InternalServerError(ctx)
		log.Println(err)
		return
	}
	Created(ctx, token)
	return
}

func (a accountHandler) SignIn(ctx *gin.Context) {
	account, err := convertAccount(ctx)
	if err != nil {
		BadRequest(ctx, err)
		log.Println(err)
		return
	}
	token, err := a.uc.SignIn(account)
	if err != nil {
		UnAuthorized(ctx)
		log.Println(err)
		return
	}
	data := map[string]interface{}{
		"token": token,
	}
	OK(ctx, data)
	return
}

func (a accountHandler) SignOut(ctx *gin.Context) {
	panic("implement me")
}

func (a accountHandler) UpdatePassword(ctx *gin.Context) {
	panic("implement me")
}

func (a accountHandler) UpdateProfile(ctx *gin.Context) {
	panic("implement me")
}

func convertAccount(ctx *gin.Context) (*usecase.AuthAccount, error) {
	var account usecase.AuthAccount
	if err := ctx.ShouldBindJSON(&account); err != nil {
		return nil, err
	}
	account.Token = ctx.GetHeader("Authorization")
	account.ID = EscapeString(account.ID)
	account.Password = EscapeString(account.Password)

	return &account, nil
}

func NewAccountHandler(repository registry.Repository, service registry.Service) AccountHandler {
	uc := usecase.NewAccount(
		repository.NewAccountRepository(),
		service.NewAccountService(),
		service.NewAuthService(),
	)

	return &accountHandler{uc}
}
