package valueobjects

import (
	"errors"
	"regexp"
	"strings"

	"github.com/LHS-Real-Estate/cim-core/internal/pkg/validator"
)

type Document struct {
	Name      string `validate:"required"`
	FilePath  string `validate:"required,filepath"`
	Extension string `validate:"required,lowercase"`
}

var (
	ErrInvalidFilePath = errors.New("the document filepath must have the file name")
)

const EXT_REGXP = "(\\.[^.]+)$"

func NewDocument(name string, filePath string) (Document, error) {
	reg := regexp.MustCompile(EXT_REGXP)
	match := reg.FindStringSubmatch(name)
	ext := ""

	if match != nil {
		ext = strings.ToLower(match[0])
	}

	document := Document{
		Name:      name,
		FilePath:  filePath,
		Extension: strings.Replace(ext, ".", "", 1),
	}

	return document, validate(document)
}

func validate(d Document) error {
	cv := validator.NewCustomValidate()
	err := cv.Validate(d)

	if !strings.HasSuffix(d.FilePath, d.Name) {
		err = errors.Join(err, ErrInvalidFilePath)
	}

	return err
}
