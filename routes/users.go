package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustaphalimar/event-booking/models"
)

func signUp(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Failed to bind json data.",
		})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Could not save user.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}
