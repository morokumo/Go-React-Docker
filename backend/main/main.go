package main

import (
	"backend/adapter/handler"
	"backend/adapter/registry"
	"backend/domain/entity"
	"backend/infra/dao"
	"backend/utility"
	"github.com/gin-gonic/gin"
)

//
//func setUpRepository(){
//	repository := registry.NewRepository()
//	db, _ := dao.ConnectDB()
//	db.AutoMigrate(&entity.Account{})
//	db.AutoMigrate(&entity.Message{})
//	db.AutoMigrate(&entity.Room{})
//}
//
//
//func setUpHandler() {
//
//
//	service := registry.NewService(repository)
//	accountHandler := handler.NewAccountHandler(repository, service)
//	messageHandler := handler.NewMessageHandler(repository, service)
//}
//

func setupRouter() *gin.Engine {
	repository := registry.NewRepository()
	db, _ := dao.ConnectDB()
	db.AutoMigrate(&entity.Account{})
	db.AutoMigrate(&entity.Message{})
	db.AutoMigrate(&entity.Room{})
	service := registry.NewService(repository)
	accountHandler := handler.NewAccountHandler(repository, service)
	messageHandler := handler.NewMessageHandler(repository, service)
	router := gin.Default()
	router.POST("/verify", func(context *gin.Context) {
		if err := accountHandler.Verify(context); err == nil {
			utility.OK(context, nil)
		}
	})
	router.POST("/signUp", func(context *gin.Context) {
		accountHandler.SignUp(context)
	})
	router.POST("/signIn", func(context *gin.Context) {
		accountHandler.SignIn(context)
	})

	r := router.Group("/api")

	r.Use(func(context *gin.Context) {
		accountHandler.Verify(context)
	})

	r.POST("/createRoom", func(c *gin.Context) {
		messageHandler.CreateRoom(c)
	})

	r.POST("/findMyRoom", func(c *gin.Context) {
		messageHandler.FindMyRooms(c)
	})

	r.POST("/findPublicRoom", func(c *gin.Context) {
		messageHandler.FindPublicRooms(c)
	})
	r.POST("/findRoomAccounts", func(c *gin.Context) {
		messageHandler.FindRoomAccount(c)
	})

	r.POST("/joinRoom", func(c *gin.Context) {
		messageHandler.JoinRoom(c)
	})

	chat := r.Group("/chat")

	chat.Use(func(context *gin.Context) {
		messageHandler.Verify(context)
	})

	chat.POST("/sendMessage", func(context *gin.Context) {
		messageHandler.SendMessage(context)
	})
	chat.POST("/getMessage", func(context *gin.Context) {
		messageHandler.GetMessage(context)
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
