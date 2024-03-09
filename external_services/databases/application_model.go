package databases

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kajiLabTeam/xr-project-application-authentication-server/utils"
	"github.com/lib/pq"
)

type Application struct {
	id        *string
	secretKey *string
	name      string
	createdAt *time.Time
	updatedAt *time.Time
	deleteAt  *time.Time
}

func NewApplication(name string) (*Application, error) {
	uuid := utils.GenerateUuid()
	ulid, err := utils.GenerateUlid()
	if err != nil {
		return nil, err
	}

	if len(name) > 50 {
		return nil, fmt.Errorf("application name is too long")
	}

	return &Application{
		id:        &ulid,
		secretKey: &uuid,
		name:      name,
		createdAt: nil,
		updatedAt: nil,
		deleteAt:  nil,
	}, nil
}

func (a *Application) GetIdOfPrivateValue() *string {
	return a.id
}

func (a *Application) GetSecretKeyOfPrivateValue() *string {
	return a.secretKey
}

func (a *Application) GetNameOfPrivateValue() string {
	return a.name
}

func (a *Application) GetCreatedAtOfPrivateValue() *time.Time {
	return a.createdAt
}

func (a *Application) GetUpdatedAtOfPrivateValue() *time.Time {
	return a.updatedAt
}

func (a *Application) GetDeleteAtOfPrivateValue() *time.Time {
	return a.deleteAt
}

func (a *Application) Insert(db *sql.DB) (*Application, error) {
	var insertedId string
	var insertedSecretKey string
	var insertedName string
	var createdAt time.Time
	var updatedAt pq.NullTime
	var deletedAt pq.NullTime

	err := db.QueryRow(
		"INSERT INTO applications (id, secret_key, name) VALUES ($1, $2, $3) RETURNING id, secret_key, name, created_at, updated_at, deleted_at",
		a.id, a.secretKey, a.name,
	).Scan(&insertedId, &insertedSecretKey, &insertedName, &createdAt, &updatedAt, &deletedAt)
	if err != nil {
		return nil, err
	}

	return &Application{
		id:        &insertedId,
		secretKey: &insertedSecretKey,
		name:      insertedName,
		createdAt: &createdAt,
		updatedAt: nil,
		deleteAt:  nil,
	}, nil
}

func FindApplicationById(db *sql.DB, id string) (*Application, error) {
	var application Application
	err := db.QueryRow("SELECT id, secret_key, name, created_at, updated_at, deleted_at FROM applications WHERE id = $1", id).Scan(
		&application.id, &application.secretKey, &application.name, &application.createdAt, &application.updatedAt, &application.deleteAt,
	)
	if err != nil {
		return nil, err
	}

	return &application, nil
}
