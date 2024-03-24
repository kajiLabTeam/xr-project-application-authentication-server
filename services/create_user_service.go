package services

import (
	"github.com/kajiLabTeam/xr-project-application-authentication-server/config"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/models"
)

type CreateUserService struct{}

func (cs *CreateUserService) Run(u *models.User) (*models.User, error) {
	conn := config.DBConnection{}
	db, err := conn.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// userのインスタンスをDBに保存
	resUser, err := u.Insert(db)
	if err != nil {
		return nil, err
	}

	return resUser, nil
}
