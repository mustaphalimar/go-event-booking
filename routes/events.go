package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustaphalimar/event-booking/models"
)

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
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	userId := ctx.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    event,
	})
}

func updateEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Not authorized to update this event.",
		})
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func deleteEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Event with id  not found!",
		})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Not authorized to delete this event.",
		})
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
