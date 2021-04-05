package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response interface {
	OK(ctx *gin.Context) error
	Created(ctx *gin.Context) error
	BadRequest(ctx *gin.Context) error
	UnAuthorized(ctx *gin.Context) error
	Forbidden(ctx *gin.Context) error
	InternalServerError(ctx *gin.Context)
}

func OK(ctx *gin.Context, data map[string]interface{}) error {
	if data == nil{
		data = map[string]interface{}{}
	}
	data["message"] = "Status OK"
	ctx.JSON(http.StatusOK, data)
	ctx.Abort()

	return nil
}

func Created(ctx *gin.Context, token string) error {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Created", "token": token})
	ctx.Abort()
	return nil
}

func BadRequest(ctx *gin.Context, err error) error {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "error": err})
	ctx.Abort()
	return nil
}
func UnAuthorized(ctx *gin.Context) error {
	ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	ctx.Abort()
	return nil
}
func Forbidden(ctx *gin.Context) error {
	ctx.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
	ctx.Abort()
	return nil
}
func InternalServerError(ctx *gin.Context) error {
	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	ctx.Abort()
	return nil
}
