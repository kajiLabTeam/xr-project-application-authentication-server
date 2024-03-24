package services

import (
	"github.com/kajiLabTeam/xr-project-application-authentication-server/config"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/models"
)

type RegisterApplicationService struct{}

func (r *RegisterApplicationService) Run(
	app *models.Application,
	rep *models.Representative,
) (*models.Application, *models.Representative, error) {
	conn := config.DBConnection{}
	db, err := conn.Connect()
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()

	// applicationのインスタンスをDBに保存
	resApp, err := app.Insert(db)
	if err != nil {
		return nil, nil, err
	}

	// representativeのインスタンスをDBに保存
	resRep, err := rep.Insert(db)
	if err != nil {
		return nil, nil, err
	}

	return resApp, resRep, nil
}
