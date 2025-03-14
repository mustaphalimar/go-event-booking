package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustaphalimar/event-booking/db"
	"github.com/mustaphalimar/event-booking/models"
)

func main() {
	db.InitDb()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", models.CreateEvent)

	server.Run(":8080") // dev env: localhost:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}
	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    event,
	})
}
