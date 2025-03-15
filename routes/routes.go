package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustaphalimar/event-booking/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	// Events routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/cancel-registration", cancelRegistration)

	// Auth routes
	server.POST("/signup", signUp)
	server.POST("/login", login)

}
