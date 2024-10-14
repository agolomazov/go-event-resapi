package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User create successfully"})
}