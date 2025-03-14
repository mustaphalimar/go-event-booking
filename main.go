package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mustaphalimar/event-booking/db"
	"github.com/mustaphalimar/event-booking/routes"
)

func main() {
	db.InitDb()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // dev env: localhost:8080
}
