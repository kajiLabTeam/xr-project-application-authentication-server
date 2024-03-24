package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/handlers/middleware"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/models"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/services"
)

type CreateUserRequest struct {
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Gender     string  `json:"gender"`
	Age        int     `json:"age"`
	Height     float64 `json:"height"`
	Weight     float64 `json:"weight"`
	Occupation string  `json:"occupation"`
	Address    string  `json:"address"`
}

type CreateUserResponse struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Gender     string  `json:"gender"`
	Age        int     `json:"age"`
	Height     float64 `json:"height"`
	Weight     float64 `json:"weight"`
	Occupation string  `json:"occupation"`
	Address    string  `json:"address"`
}

func CreateUserHandler(r *gin.Engine) {
	r.POST("api/user/create", middleware.AuthApplicationMiddleware(), func(c *gin.Context) {
		cus := services.CreateUserService{}

		applicationId, exists := c.Get("applicationId")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "AuthResult not found"})
			return
		}

		applicationIdStr, ok := applicationId.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "AuthResult type assertion failed"})
			return
		}

		var req CreateUserRequest

		// リクエストのバリデーション
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u, err := models.NewUser(
			req.Name,
			req.Email,
			req.Gender,
			req.Age,
			req.Height,
			req.Weight,
			req.Occupation,
			req.Address,
			nil,
			nil,
			nil,
			applicationIdStr,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		resUser, err := cus.Run(u)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		res := CreateUserResponse{
			Id:         resUser.GetApplicationIdOfPrivateValue(),
			Name:       resUser.GetNameOfPrivateValue(),
			Email:      resUser.GetMailOfPrivateValue(),
			Gender:     resUser.GetGenderOfPrivateValue(),
			Age:        resUser.GetAgeOfPrivateValue(),
			Height:     resUser.GetHeightOfPrivateValue(),
			Weight:     resUser.GetHeightOfPrivateValue(),
			Occupation: resUser.GetOccupationOfPrivateValue(),
			Address:    resUser.GetAddressOfPrivateValue(),
		}

		// レスポンスを返却
		c.JSON(http.StatusCreated, res)
	})
}
