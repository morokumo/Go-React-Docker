package handler

import (
	"backend/adapter/registry"
	"backend/domain/entity"
	"backend/usecase"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	"io"
	"io/ioutil"
	"log"
)

type MessageHandler interface {
	FindMyRooms(ctx *gin.Context)
	FindPublicRooms(ctx *gin.Context)
	CreateRoom(ctx *gin.Context)
	DeleteRoom(ctx *gin.Context) error
	JoinRoom(ctx *gin.Context)
	LeaveRoom(ctx *gin.Context) error
	UpdateRoom(ctx *gin.Context) error
	SendMessage(ctx *gin.Context)
	GetMessage(ctx *gin.Context)
	Verify(ctx *gin.Context)
}
type message struct {
	uc usecase.Message
}
type jsonRoom struct {
	Name    string `json:"room_name"`
	Info    string `json:"info"`
	Private string `json:"private"`
}

func (m message) Verify(ctx *gin.Context) {
	bodyCopy := new(bytes.Buffer)
	_, err := io.Copy(bodyCopy, ctx.Request.Body)
	if err != nil {
		BadRequest(ctx, err)
	}
	bodyData := bodyCopy.Bytes()
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
	sendContent, err := convertSendContent(ctx)
	if err != nil {
		BadRequest(ctx, err)
	}
	if err := m.uc.Verify(sendContent); err != nil {
		UnAuthorized(ctx)
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
}

func (m message) FindMyRooms(ctx *gin.Context) {
	account, err := readAccount(ctx)
	if err != nil {
		UnAuthorized(ctx)
		return
	}
	rooms, err := m.uc.FindMyRooms(account)
	if err != nil {
		InternalServerError(ctx)
		return
	}
	data := map[string]interface{}{
		"rooms": rooms,
	}
	OK(ctx, data)

}
func (m message) FindPublicRooms(ctx *gin.Context) {
	account, err := readAccount(ctx)
	if err != nil {
		UnAuthorized(ctx)
		return
	}
	rooms, err := m.uc.FindPublicRooms(account)
	if err != nil {
		InternalServerError(ctx)
		return
	}
	data := map[string]interface{}{
		"rooms": rooms,
	}
	OK(ctx, data)

}

func (m message) CreateRoom(ctx *gin.Context) {
	room, err := convertRoom(ctx)
	if err != nil {
		BadRequest(ctx, err)
		return
	}
	err = m.uc.CreateRoom(room)
	if err != nil {
		InternalServerError(ctx)
		return
	}
	Created(ctx, "")
}

func (m message) DeleteRoom(ctx *gin.Context) error {
	panic("implement me")
}

func (m message) JoinRoom(ctx *gin.Context) {
	sendContent, err := convertSendContent(ctx)
	log.Println(sendContent)
	if err != nil {
		BadRequest(ctx, err)
		return
	}
	m.uc.JoinRoom(sendContent)
	if err != nil {
		InternalServerError(ctx)
		return
	}
	OK(ctx, nil)
}

func (m message) LeaveRoom(ctx *gin.Context) error {
	panic("implement me")
}

func (m message) UpdateRoom(ctx *gin.Context) error {
	panic("implement me")
}

func (m message) SendMessage(ctx *gin.Context) {
	sendContent, err := convertSendContent(ctx)
	if err != nil {
		BadRequest(ctx, err)
		return
	}
	msg, err := m.uc.SendMessage(sendContent)
	if err != nil {
		InternalServerError(ctx)
		return
	}
	data := map[string]interface{}{
		"chat_message": msg,
	}
	OK(ctx, data)
}

func (m message) GetMessage(ctx *gin.Context) {
	sendContent, err := convertSendContent(ctx)
	if err != nil {
		BadRequest(ctx, err)
		return
	}
	messages, err := m.uc.GetMessageByRoom(sendContent)
	if err != nil {
		InternalServerError(ctx)
		return
	}
	data := map[string]interface{}{
		"chat_messages": messages,
	}
	OK(ctx, data)
}

func readAccount(ctx *gin.Context) (*entity.Account, error) {
	accountI, ok := ctx.Get("account")
	if !ok {
		return nil, errors.New("not found")
	}
	account := accountI.(*entity.Account)
	return account, nil
}

func convertSendContent(ctx *gin.Context) (*usecase.SendContent, error) {
	account, err := readAccount(ctx)
	if err != nil {
		panic(err)
	}
	var json usecase.SendContent
	err = ctx.ShouldBindJSON(&json)
	if err != nil {
		return nil, err
	}
	json.Account = account
	json.RoomId = EscapeString(json.RoomId)
	json.MessageText = EscapeString(json.MessageText)

	return &json, err
}
func convertRoom(ctx *gin.Context) (*entity.Room, error) {
	account, err := readAccount(ctx)
	if err != nil {
		panic(err)
	}
	var json jsonRoom
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
			Name:     EscapeString(json.Name),
			Info:     EscapeString(json.Info),
			Accounts: accounts,
			Private:  private},
		err
}

func NewMessageHandler(repository registry.Repository, service registry.Service) MessageHandler {
	uc := usecase.NewMessage(repository.NewMessageRepository(), repository.NewRoomRepository(), service.NewMessageService())
	return &message{uc}
}
