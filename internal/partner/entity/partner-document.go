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

type PartnerDocument struct {
	ID          string                `validate:"required,uuid"`
	PartnerID   string                `validate:"required,uuid"`
	Title       string                `validate:"required,min=3"`
	File        valueobjects.Document `validate:"required"`
	LastUpdated time.Time             `validate:"required,gtefield=CreatedAt"`
	CreatedAt   time.Time             `validate:"required,ltefield=LastUpdated"`
}

func NewDocument(id string, partnerID string, title string, filePath string, fileExtension string,
	lastUpdated time.Time, createdAt time.Time) (PartnerDocument, error) {

	if id == "" {
		id = uuid.New().String()
	}

	if filePath == "" {
		filePath = path.Join(PartnerRootPath(partnerID), DocumentEncryptedName(id, fileExtension))
	}

	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	if lastUpdated.IsZero() {
		lastUpdated = time.Now()
	}

	docFile := valueobjects.Document{FilePath: filePath, Extension: fileExtension}

	document := PartnerDocument{
		ID:          id,
		PartnerID:   partnerID,
		Title:       title,
		File:        docFile,
		LastUpdated: lastUpdated,
		CreatedAt:   createdAt,
	}

	return document, validatePartnerDoc(document)
}

func validatePartnerDoc(pd PartnerDocument) error {
	cv := validator.NewCustomValidate()
	err := cv.Validate(pd)
	return err
}

func DocumentEncryptedName(docID string, extension string) string {
	hash := md5.Sum([]byte(docID))
	return hex.EncodeToString(hash[:]) + "." + extension
}
