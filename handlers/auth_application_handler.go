package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/handlers/middleware"
)

func AuthApplicationHandler(r *gin.Engine) {
	r.GET("api/application/auth", middleware.AuthApplicationMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
}
