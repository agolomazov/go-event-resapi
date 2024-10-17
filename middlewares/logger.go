package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")

	if userId != 0 {
		fmt.Println("Request was exec user with id", userId)
	} else {
		fmt.Println("Request was exec unauthorized user");
	}

	ctx.Next()
}