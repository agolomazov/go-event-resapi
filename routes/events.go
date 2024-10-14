package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not send request"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func getEventById(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event by id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func createEvent(ctx *gin.Context) {
	var event models.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request error"})
		return
	}

	event.ID = 1
	event.UserId = 1
	newEvent, err := event.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request error"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": newEvent})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event by id"})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not update event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event by id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Could not remove event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event was delete"})
}