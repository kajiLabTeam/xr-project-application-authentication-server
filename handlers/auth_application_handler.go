package handlers

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/external_services"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/external_services/databases"
)

func AuthApplicationHandler(r *gin.Engine) {
	r.GET("api/application/auth", func(c *gin.Context) {
		// リクエストヘッダーから特定のキーの値を取得
		headerValue := c.GetHeader("Authorization")

		authParts := strings.Fields(headerValue)
		if len(authParts) != 2 || authParts[0] != "Basic" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// base64エンコードされた文字列をデコード
		decodedBytes, err := base64.StdEncoding.DecodeString(authParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// デコードされた文字列を ":" で分割してユーザー名とパスワードを取得
		credentials := strings.SplitN(string(decodedBytes), ":", 2)
		if len(credentials) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		appId := credentials[0]
		secretKey := credentials[1]

		// DBとのコネクションを確立
		conn := external_services.DBConnection{}
		db, err := conn.Connect()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		defer db.Close()

		// アプリケーションIDを用いてアプリケーションを検索
		application, err := databases.FindApplicationById(db, appId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// シークレットキーを検証
		if *application.GetSecretKeyOfPrivateValue() != secretKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// 認証成功
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
}
