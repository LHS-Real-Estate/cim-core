package entity

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/LHS-Real-Estate/cim-core/internal/pkg/validator"
	"github.com/google/uuid"
)

type Partner struct {
	ID        string    `validate:"required,uuid"`
	CompanyID string    `validate:"required,uuid"`
	Name      string    `validate:"required,min=2"`
	Surname   string    `validate:"omitempty,min=3"`
	IsActive  bool      `validate:"-"`
	CreatedAt time.Time `validate:"required"`
}

func NewPartner(id string, companyID string, name string, surname string, isActive bool, createdAt time.Time) (Partner, error) {

	if id == "" {
		id = uuid.New().String()
	}

	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	p := Partner{
		ID:        id,
		CompanyID: companyID,
		Name:      name,
		Surname:   surname,
		IsActive:  isActive,
		CreatedAt: createdAt,
	}

	return p, validatePartner(p)
}

func validatePartner(c Partner) error {
	cv := validator.NewCustomValidate()
	err := cv.Validate(c)

	return err
}

func PartnerRootPath(partnerID string) string {
	hash := md5.Sum([]byte(partnerID))
	return "Partner-" + hex.EncodeToString(hash[:])
}
