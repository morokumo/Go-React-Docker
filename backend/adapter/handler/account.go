package handler

import (
	"backend/adapter/registry"
	"backend/usecase"
	"backend/utility"
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
		utility.BadRequest(ctx, err)
		return err
	}
	bodyData := bodyCopy.Bytes()
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
	authAccount, err := convertAccount(ctx)
	if err != nil {
		utility.BadRequest(ctx, err)
		return err
	}

	account, err := a.uc.Verify(authAccount)

	if account == nil || account.ID == "" {
		utility.UnAuthorized(ctx)
		return errors.New(" Account not found")
	}

	if err != nil {
		utility.UnAuthorized(ctx)
		return err
	}

	ctx.Set("account", account)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
	return nil
}

func (a accountHandler) SignUp(ctx *gin.Context) {
	account, err := convertAccount(ctx)

	if err != nil {
		utility.BadRequest(ctx, err)
		return
	}

	if account.ID == "" || account.Password == "" {
		utility.BadRequest(ctx, errors.New(" ID and password must be at least one character."))
	}

	err = a.uc.SignUp(*account)
	if err != nil {

		return
	}
	token, err := a.uc.SignIn(account)
	if err != nil {
		utility.InternalServerError(ctx)
		return
	}
	utility.Created(ctx, token)
	return
}

func (a accountHandler) SignIn(ctx *gin.Context) {
	account, err := convertAccount(ctx)
	if err != nil {
		utility.BadRequest(ctx, err)
		log.Println(err)
		return
	}
	token, err := a.uc.SignIn(account)
	if err != nil {
		utility.UnAuthorized(ctx)
		log.Println(err)
		return
	}
	data := map[string]interface{}{
		"token": token,
	}
	utility.OK(ctx, data)
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
		return nil, errors.New(" Request invalid.")
	}
	account.Token = ctx.GetHeader("Authorization")
	account.ID = utility.EscapeString(account.ID)
	account.Password = utility.EscapeString(account.Password)
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
