package valueobjects_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/LHS-Real-Estate/cim-core/internal/pkg/entity/valueobjects"
	"github.com/stretchr/testify/assert"
)

func TestDocument_NewDocument(t *testing.T) {
	type input struct {
		name     string
		filePath string
	}

	type output struct {
		name      string
		filePath  string
		extension string
	}

	type testCase struct {
		test           string
		input          input
		expectedOutput output
		expectedError  error
	}

	testsTable := []testCase{
		{
			test:           "Empty document name validation",
			input:          input{name: "", filePath: "path/to/file/"},
			expectedOutput: output{filePath: "path/to/file/"},
			expectedError:  errors.New("invalid fields: Document.Name: \"\", Document.Extension: \"\""),
		},
		{
			test:           "Empty document filePath validation",
			input:          input{name: "Document Name.txt", filePath: ""},
			expectedOutput: output{name: "Document Name.txt", extension: "txt"},
			expectedError:  errors.New("invalid fields: Document.FilePath: \"\""),
		},
		{
			test:           "Document without extension validation",
			input:          input{name: "Document_without_extension", filePath: "path/to/file/Document_without_extension"},
			expectedOutput: output{name: "Document_without_extension", filePath: "path/to/file/Document_without_extension"},
			expectedError:  errors.New("invalid fields: Document.Extension: \"\""),
		},
		{
			test:           "Document filePath without name validation",
			input:          input{name: "Document Name.txt", filePath: "path/to/"},
			expectedOutput: output{name: "Document Name.txt", filePath: "path/to/", extension: "txt"},
			expectedError:  errors.New("the document filepath must have the file name"),
		},
		{
			test:           "Valid Document instance",
			input:          input{name: "Document name.txt", filePath: "path/to/Document name.txt"},
			expectedOutput: output{"Document name.txt", "path/to/Document name.txt", "txt"},
			expectedError:  nil,
		},
	}

	for _, tc := range testsTable {
		fmt.Printf("Test case: %s", tc.test)
		doc, err := valueobjects.NewDocument(tc.input.name, tc.input.filePath)

		assert.Equal(t, doc.Name, tc.expectedOutput.name)
		assert.Equal(t, doc.FilePath, tc.expectedOutput.filePath)
		assert.Equal(t, doc.Extension, tc.expectedOutput.extension)

		if tc.expectedError != nil {
			assert.Error(t, err, tc.expectedError)
			continue
		}

		assert.Nil(t, err)
	}
}
