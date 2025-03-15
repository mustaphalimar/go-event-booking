package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustaphalimar/event-booking/utils"
)

func Authenticate(ctx *gin.Context) {
	tokenHeader := ctx.Request.Header.Get("Authentication")

	if tokenHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized, token not provided!",
		})
		return
	}

	userId, err := utils.VerifyToken(tokenHeader)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	ctx.Set("userId", userId)
	ctx.Next()
}
