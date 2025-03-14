package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustaphalimar/event-booking/db"
	"github.com/mustaphalimar/event-booking/models"
)

func main() {
	db.InitDb()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080") // dev env: localhost:8080
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    event,
	})
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}
	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    event,
	})
}
