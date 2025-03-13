package models

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Datetime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (event Event) Save() {
	// later:
	events = append(events, event)
}

func GetEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    events,
	})
}

func CreateEvent(context *gin.Context) {
	var event Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"err":     err,
		})
		return
	}
	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    event,
	})
}
