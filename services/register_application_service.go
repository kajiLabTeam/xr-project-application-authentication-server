package services

import (
	"github.com/kajiLabTeam/xr-project-application-authentication-server/external_services"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/external_services/databases"
)

type RegisterApplicationService struct{}

func (r *RegisterApplicationService) Run(app *databases.Application, rep *databases.Representative) (*databases.Application, *databases.Representative, error) {
	conn := external_services.DBConnection{}
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
