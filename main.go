package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mustaphalimar/event-booking/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", models.GetEvents)
	server.POST("/events", models.CreateEvent)

	server.Run(":8080") // dev env: localhost:8080
}
