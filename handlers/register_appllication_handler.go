package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/external_services/databases"

	"github.com/kajiLabTeam/xr-project-application-authentication-server/services"
)

type AuthApplicationRequest struct {
	ApplicationName    string `json:"applicationName"`
	OrganizationName   string `json:"organizationName"`
	RepresentativeName string `json:"representativeName"`
	PhoneNumber        string `json:"phoneNumber"`
	Email              string `json:"email"`
	Address            string `json:"address"`
}

type AuthApplicationResponse struct {
	ApplicationName    string `json:"applicationName"`
	OrganizationName   string `json:"organizationName"`
	RepresentativeName string `json:"representativeName"`
	PhoneNumber        string `json:"phoneNumber"`
	Email              string `json:"email"`
	Address            string `json:"address"`
	AccessKey          string `json:"accessKey"`
	SecretKey          string `json:"secretKey"`
}

func RegisterApplicationHandler(r *gin.Engine) {
	r.POST("api/application/create", func(c *gin.Context) {
		ras := services.RegisterApplicationService{}

		var req AuthApplicationRequest

		// リクエストのバリデーション
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// applicationのインスタンスを生成
		app, err := databases.NewApplication(req.ApplicationName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// representativeのインスタンスを生成
		rep, err := databases.NewRepresentative(
			req.RepresentativeName,
			req.OrganizationName,
			req.Email,
			req.PhoneNumber,
			req.Address,
			*app.GetIdOfPrivateValue(),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		resApp, resRep, err := ras.Run(app, rep)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var res = AuthApplicationResponse{
			ApplicationName:    resApp.GetNameOfPrivateValue(),
			OrganizationName:   resRep.GetCorporateNameOfPrivateValue(),
			RepresentativeName: resRep.GetNameOfPrivateValue(),
			PhoneNumber:        resRep.GetPhoneNumberOfPrivateValue(),
			Email:              resRep.GetMailOfPrivateValue(),
			Address:            resRep.GetAddressOfPrivateValue(),
			AccessKey:          *resApp.GetIdOfPrivateValue(),
			SecretKey:          *resApp.GetSecretKeyOfPrivateValue(),
		}

		// レスポンスを返却
		c.JSON(http.StatusCreated, res)
	})
}
