package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/config"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/handlers/middleware"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/models"
)

type AuthUserRequest struct {
	Id string `json:"id"`
}

func AuthUserHandler(r *gin.Engine) {
	r.POST("api/user/auth",
		middleware.AuthApplicationMiddleware(),
		func(c *gin.Context) {
			var req AuthUserRequest

			// リクエストのバリデーション
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// DBとのコネクションを確立
			conn := config.DBConnection{}
			db, err := conn.Connect()
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user"})
				return
			}
			defer db.Close()

			// UserIDを用いてユーザを検索
			_, err = models.FindUserById(db, req.Id)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "ok"})
		})
}
