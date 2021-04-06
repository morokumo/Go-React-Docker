package handler

import (
	"backend/adapter/registry"
	"backend/domain/entity"
	"backend/usecase"
	"backend/utility"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	"io"
	"io/ioutil"
)

type MessageHandler interface {
	Verify(ctx *gin.Context)
	FindMyRooms(ctx *gin.Context)
	FindPublicRooms(ctx *gin.Context)
	FindRoomAccount(ctx *gin.Context)
	CreateRoom(ctx *gin.Context)
	DeleteRoom(ctx *gin.Context) error
	JoinRoom(ctx *gin.Context)
	LeaveRoom(ctx *gin.Context) error
	UpdateRoom(ctx *gin.Context) error
	SendMessage(ctx *gin.Context)
	GetMessage(ctx *gin.Context)
}
type message struct {
	uc usecase.Message
}

func (m message) Verify(ctx *gin.Context) {
	bodyCopy := new(bytes.Buffer)
	_, err := io.Copy(bodyCopy, ctx.Request.Body)
	if err != nil {
		utility.BadRequest(ctx, err)
	}
	bodyData := bodyCopy.Bytes()
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))

	sendContent, err := convertRequestMessage(ctx)

	if err != nil {
		utility.BadRequest(ctx, err)
	}

	if err := m.uc.Verify(sendContent); err != nil {
		utility.UnAuthorized(ctx)
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
}

func (m message) FindMyRooms(ctx *gin.Context) {
	account, err := readAccount(ctx)
	if err != nil {
		utility.UnAuthorized(ctx)
		return
	}
	rooms, err := m.uc.FindMyRooms(account)
	if err != nil {
		utility.InternalServerError(ctx)
		return
	}
	data := map[string]interface{}{
		"rooms": rooms,
	}
	utility.OK(ctx, data)

}
func (m message) FindPublicRooms(ctx *gin.Context) {
	account, err := readAccount(ctx)
	if err != nil {
		utility.UnAuthorized(ctx)
		return
	}
	rooms, err := m.uc.FindPublicRooms(account)
	if err != nil {
		utility.InternalServerError(ctx)
		return
	}
	data := map[string]interface{}{
		"rooms": rooms,
	}
	utility.OK(ctx, data)

}

func (m message) CreateRoom(ctx *gin.Context) {
	room, err := convertRoom(ctx)
	if err != nil {
		utility.BadRequest(ctx, err)
		return
	}
	err = m.uc.CreateRoom(room)
	if err != nil {
		utility.InternalServerError(ctx)
		return
	}
	utility.Created(ctx, "")
}
func (m message) FindRoomAccount(ctx *gin.Context) {
	sendContent, err := convertRequestMessage(ctx)
	if err != nil {
		utility.BadRequest(ctx, err)
		return
	}
	accounts, err := m.uc.FindRoomAccounts(sendContent)
	if err != nil {
		utility.InternalServerError(ctx)
		return
	}

	data := map[string]interface{}{
		"accounts": accounts,
	}
	utility.OK(ctx, data)
}

func (m message) DeleteRoom(ctx *gin.Context) error {
	panic("implement me")
}

func (m message) JoinRoom(ctx *gin.Context) {
	sendContent, err := convertRequestMessage(ctx)
	if err != nil {
		utility.BadRequest(ctx, err)
		return
	}
	m.uc.JoinRoom(sendContent)
	if err != nil {
		utility.InternalServerError(ctx)
		return
	}
	utility.OK(ctx, nil)
}

func (m message) LeaveRoom(ctx *gin.Context) error {
	panic("implement me")
}

func (m message) UpdateRoom(ctx *gin.Context) error {
	panic("implement me")
}

func (m message) SendMessage(ctx *gin.Context) {
	sendContent, err := convertRequestMessage(ctx)
	if err != nil {
		utility.BadRequest(ctx, err)
		return
	}
	msg, err := m.uc.SendMessage(sendContent)
	if err != nil {
		utility.InternalServerError(ctx)
		return
	}
	data := map[string]interface{}{
		"chat_message": msg,
	}
	utility.OK(ctx, data)
}

func (m message) GetMessage(ctx *gin.Context) {
	sendContent, err := convertRequestMessage(ctx)
	if err != nil {
		utility.BadRequest(ctx, err)
		return
	}
	messages, err := m.uc.GetMessageByRoom(sendContent)
	if err != nil {
		utility.InternalServerError(ctx)
		return
	}
	data := map[string]interface{}{
		"chat_messages": messages,
	}
	utility.OK(ctx, data)
}

func readAccount(ctx *gin.Context) (*entity.Account, error) {
	accountI, ok := ctx.Get("account")
	if !ok {
		return nil, errors.New("not found")
	}
	account := accountI.(*entity.Account)
	return account, nil
}

func convertRequestMessage(ctx *gin.Context) (*usecase.RequestMessage, error) {
	account, err := readAccount(ctx)
	if err != nil {
		panic(err)
	}
	var requestMessage usecase.RequestMessage
	err = ctx.ShouldBindJSON(&requestMessage)
	if err != nil {
		return nil, err
	}
	requestMessage.Account = account
	requestMessage.RoomId = utility.EscapeString(requestMessage.RoomId)
	requestMessage.MessageText = utility.EscapeString(requestMessage.MessageText)

	return &requestMessage, err
}
func convertRoom(ctx *gin.Context) (*entity.Room, error) {
	account, err := readAccount(ctx)
	if err != nil {
		return nil, err
	}
	var json usecase.RequestRoom
	err = ctx.ShouldBindJSON(&json)
	if err != nil {
		return nil, err
	}
	if json.Name == "" {
		return nil, errors.New("room name is not set.")
	}
	var accounts []entity.Account
	accounts = append(accounts, *account)
	private := true
	if json.Private == "public" {
		private = false
	}

	return &entity.Room{
			ID:       uuid2.NewString(),
			Name:     utility.EscapeString(json.Name),
			Info:     utility.EscapeString(json.Info),
			Accounts: accounts,
			Private:  private},
		err
}

func NewMessageHandler(repository registry.Repository, service registry.Service) MessageHandler {
	uc := usecase.NewMessage(repository.NewMessageRepository(), repository.NewRoomRepository(), service.NewMessageService())
	return &message{uc}
}
