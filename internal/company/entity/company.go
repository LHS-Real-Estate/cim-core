package entity

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/LHS-Real-Estate/cim-core/internal/pkg/validator"
	"github.com/google/uuid"
)

type Company struct {
	ID                    string    `validate:"required,uuid"`
	EIN                   string    `validate:"required"`
	Name                  string    `validate:"required,min=3"`
	FullName              string    `validate:"required,min=3"`
	MunicipalRegistration string    `validate:""`
	StateRegistration     string    `validate:""`
	CreatedAt             time.Time `validate:"required"`
}

func NewCompany(id string, ein string, name string, fullName string, municipalReg string, stateReg string,
	createdAt time.Time) (Company, error) {

	if id == "" {
		id = uuid.New().String()
	}

	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	company := Company{
		ID:                    id,
		EIN:                   ein,
		Name:                  name,
		FullName:              fullName,
		MunicipalRegistration: municipalReg,
		StateRegistration:     stateReg,
		CreatedAt:             createdAt,
	}

	return company, validateCompany(company)
}

func validateCompany(c Company) error {
	cv := validator.NewCustomValidate()
	err := cv.Validate(c)

	return err
}

func CompanyRootPath(companyID string) string {
	hash := md5.Sum([]byte(companyID))
	return "Company-" + hex.EncodeToString(hash[:])
}
