package utility

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OK(ctx *gin.Context, data map[string]interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	data["message"] = "Status OK"
	ctx.JSON(http.StatusOK, data)
	ctx.Abort()
}

func Created(ctx *gin.Context, token string) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Created", "token": token})
	ctx.Abort()
}

func BadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "error": err})
	ctx.Abort()
}
func UnAuthorized(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	ctx.Abort()
}
func Forbidden(ctx *gin.Context) {
	ctx.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
	ctx.Abort()
}
func InternalServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	ctx.Abort()
}
