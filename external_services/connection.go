package external_services

import (
	"database/sql"
	"fmt"

	"github.com/kajiLabTeam/xr-project-application-authentication-server/config"

	_ "github.com/lib/pq"
)

type DBConnection struct{}

func (dbc *DBConnection) Connect() (*sql.DB, error) {
	env := config.NewPostgresEnv()

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.GetHostOfPrivateValue(),
		env.GetPortOfPrivateValue(),
		env.GetUserOfPrivateValue(),
		env.GetPasswordOfPrivateValue(),
		env.GetDatabaseOfPrivateValue(),
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
