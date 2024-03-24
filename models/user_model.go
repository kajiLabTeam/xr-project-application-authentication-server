package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kajiLabTeam/xr-project-application-authentication-server/utils"
)

type User struct {
	id            *string
	name          string
	mail          string
	gender        string
	age           int
	height        float64
	weight        float64
	occupation    string
	address       string
	createAt      *time.Time
	updateAt      *time.Time
	deleteAt      *time.Time
	applicationId string
}

func NewUser(
	name string,
	mail string,
	gender string,
	age int,
	height float64,
	weight float64,
	occupation string,
	address string,
	createdAt *time.Time,
	updateAt *time.Time,
	deleteAt *time.Time,
	applicationId string,
) (*User, error) {
	ulid, err := utils.GenerateUlid()
	if err != nil {
		return nil, err
	}

	if len(name) > 50 {
		return nil, fmt.Errorf("name is too long")
	}

	if !utils.ValidateEmail(&mail) {
		return nil, fmt.Errorf("invalid email")
	}

	return &User{
		id:            &ulid,
		name:          name,
		mail:          mail,
		gender:        gender,
		age:           age,
		height:        height,
		weight:        weight,
		occupation:    occupation,
		address:       address,
		createAt:      createdAt,
		updateAt:      updateAt,
		deleteAt:      deleteAt,
		applicationId: applicationId,
	}, nil
}

func (u *User) GetIdOfPrivateValue() *string {
	return u.id
}

func (u *User) GetNameOfPrivateValue() string {
	return u.name
}

func (u *User) GetMailOfPrivateValue() string {
	return u.mail
}

func (u *User) GetGenderOfPrivateValue() string {
	return u.gender
}

func (u *User) GetAgeOfPrivateValue() int {
	return u.age
}

func (u *User) GetHeightOfPrivateValue() float64 {
	return u.height
}

func (u *User) GetWeightOfPrivateValue() float64 {
	return u.weight
}

func (u *User) GetOccupationOfPrivateValue() string {
	return u.occupation
}

func (u *User) GetAddressOfPrivateValue() string {
	return u.address
}

func (u *User) GetCreatedAtOfPrivateValue() *time.Time {
	return u.createAt
}

func (u *User) GetUpdatedAtOfPrivateValue() *time.Time {
	return u.updateAt
}

func (u *User) GetDeleteAtOfPrivateValue() *time.Time {
	return u.deleteAt
}

func (u *User) GetApplicationIdOfPrivateValue() string {
	return u.applicationId
}

func (u *User) Insert(db *sql.DB) (*User, error) {
	var insertedId string
	var insertedName string
	var insertedMail string
	var insertedGender string
	var insertedAge int
	var insertedHeight float64
	var insertedWeight float64
	var insertedAddress string
	var insertedOccupation string
	var insertedCreateAt time.Time
	var insertedApplicationId string

	err := db.QueryRow(
		"INSERT INTO users (id, name, mail, gender, age, height, weight, address, occupation, application_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, name, mail, gender, age, height, weight, occupation, address, created_at, application_id",
		u.id,
		u.name,
		u.mail,
		u.gender,
		u.age,
		u.height,
		u.weight,
		u.address,
		u.occupation,
		u.applicationId,
	).Scan(
		&insertedId,
		&insertedName,
		&insertedMail,
		&insertedGender,
		&insertedAge,
		&insertedHeight,
		&insertedWeight,
		&insertedAddress,
		&insertedOccupation,
		&insertedCreateAt,
		&insertedApplicationId,
	)
	if err != nil {
		return nil, err
	}

	return &User{
		id:            &insertedId,
		name:          insertedName,
		mail:          insertedMail,
		gender:        insertedGender,
		age:           insertedAge,
		height:        insertedHeight,
		weight:        insertedWeight,
		address:       insertedAddress,
		occupation:    insertedOccupation,
		createAt:      &insertedCreateAt,
		updateAt:      nil,
		deleteAt:      nil,
		applicationId: insertedApplicationId,
	}, nil
}

func FindUserById(db *sql.DB, id string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id FROM users WHERE id = $1", id).Scan(
		&user.id,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}
