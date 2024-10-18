package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventId"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Could not fetch event"})
	}

	err = event.Register(userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not create registration for user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration create successfully"})
}

func cancelRegistration(ctx *gin.Context) {}