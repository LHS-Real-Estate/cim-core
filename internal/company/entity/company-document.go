package entity

import (
	"crypto/md5"
	"encoding/hex"
	"path"
	"time"

	"github.com/LHS-Real-Estate/cim-core/internal/pkg/entity/valueobjects"
	"github.com/LHS-Real-Estate/cim-core/internal/pkg/validator"
	"github.com/google/uuid"
)

type CompanyDocument struct {
	ID          string                `validate:"required,uuid"`
	CompanyID   string                `validate:"required,uuid"`
	Title       string                `validate:"required,min=3"`
	File        valueobjects.Document `validate:"required"`
	LastUpdated time.Time             `validate:"required,gtefield=CreatedAt"`
	CreatedAt   time.Time             `validate:"required,ltefield=LastUpdated"`
}

func NewDocument(id string, companyID string, title string, filePath string, fileExtension string,
	lastUpdated time.Time, createdAt time.Time) (CompanyDocument, error) {

	if id == "" {
		id = uuid.New().String()
	}

	if filePath == "" {
		filePath = path.Join(CompanyRootPath(companyID), DocumentEncryptedName(id, fileExtension))
	}

	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	if lastUpdated.IsZero() {
		lastUpdated = time.Now()
	}

	docFile := valueobjects.Document{FilePath: filePath, Extension: fileExtension}

	document := CompanyDocument{
		ID:          id,
		CompanyID:   companyID,
		Title:       title,
		File:        docFile,
		LastUpdated: lastUpdated,
		CreatedAt:   createdAt,
	}

	return document, validateCompanyDoc(document)
}

func validateCompanyDoc(cd CompanyDocument) error {
	cv := validator.NewCustomValidate()
	err := cv.Validate(cd)
	return err
}

func DocumentEncryptedName(docID string, extension string) string {
	hash := md5.Sum([]byte(docID))
	return hex.EncodeToString(hash[:]) + "." + extension
}
