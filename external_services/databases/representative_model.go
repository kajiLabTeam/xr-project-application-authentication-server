package databases

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kajiLabTeam/xr-project-application-authentication-server/utils"
	"github.com/lib/pq"
)

type Representative struct {
	id            *string
	name          string
	corporateName string
	mail          string
	phoneNumber   string
	address       string
	createdAt     *time.Time
	updatedAt     *time.Time
	deleteAt      *time.Time
	applicationId *string
}

func NewRepresentative(
	name string,
	corporateName string,
	mail string,
	phoneNumber string,
	address string,
	applicationId string,
) (*Representative, error) {
	ulid, err := utils.GenerateUlid()
	if err != nil {
		return nil, err
	}

	if len(name) > 50 {
		return nil, fmt.Errorf("name is too long")
	}

	if len(corporateName) > 50 {
		return nil, fmt.Errorf("corporate name is too long")
	}

	if !utils.ValidateEmail(&mail) {
		return nil, fmt.Errorf("invalid email")
	}

	if !utils.ValidatePhoneNumber(&phoneNumber) {
		return nil, fmt.Errorf("invalid phone number")
	}

	if len(address) > 255 {
		return nil, fmt.Errorf("address is too long")
	}

	return &Representative{
		id:            &ulid,
		name:          name,
		corporateName: corporateName,
		mail:          mail,
		phoneNumber:   phoneNumber,
		address:       address,
		createdAt:     nil,
		updatedAt:     nil,
		deleteAt:      nil,
		applicationId: &applicationId,
	}, nil
}

func (r *Representative) GetIdOfPrivateValue() *string {
	return r.id
}

func (r *Representative) GetNameOfPrivateValue() string {
	return r.name
}

func (r *Representative) GetCorporateNameOfPrivateValue() string {
	return r.corporateName
}

func (r *Representative) GetMailOfPrivateValue() string {
	return r.mail
}

func (r *Representative) GetPhoneNumberOfPrivateValue() string {
	return r.phoneNumber
}

func (r *Representative) GetAddressOfPrivateValue() string {
	return r.address
}

func (r *Representative) GetCreatedAtOfPrivateValue() *time.Time {
	return r.createdAt
}

func (r *Representative) GetUpdatedAtOfPrivateValue() *time.Time {
	return r.updatedAt
}

func (r *Representative) GetDeleteAtOfPrivateValue() *time.Time {
	return r.deleteAt
}

func (r *Representative) GetApplicationIdOfPrivateValue() *string {
	return r.applicationId
}

func (r *Representative) Insert(db *sql.DB) (*Representative, error) {
	var insertedID string
	var insertedName string
	var insertedCorporateName string
	var insertedMail string
	var insertedPhoneNumber string
	var insertedAddress string
	var createdAt time.Time
	var updatedAt pq.NullTime
	var deletedAt pq.NullTime
	var applicationID string

	err := db.QueryRow(
		"INSERT INTO representatives (id, name, corporate_name, mail, phone_number, address, application_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, corporate_name, mail, phone_number, address, created_at, updated_at, deleted_at, application_id",
		r.id, r.name, r.corporateName, r.mail, r.phoneNumber, r.address, r.applicationId,
	).Scan(&insertedID, &insertedName, &insertedCorporateName, &insertedMail, &insertedPhoneNumber, &insertedAddress, &createdAt, &updatedAt, &deletedAt, &applicationID)
	if err != nil {
		return nil, err
	}

	return &Representative{
		id:            &insertedID,
		name:          insertedName,
		corporateName: insertedCorporateName,
		mail:          insertedMail,
		phoneNumber:   insertedPhoneNumber,
		address:       insertedAddress,
		createdAt:     &createdAt,
		updatedAt:     &updatedAt.Time,
		deleteAt:      &deletedAt.Time,
		applicationId: &applicationID,
	}, nil
}
