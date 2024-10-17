package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)
	authenticated.Use(middlewares.Logger)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}