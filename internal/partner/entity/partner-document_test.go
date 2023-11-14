package entity_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/LHS-Real-Estate/cim-core/internal/partner/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPartnerDocument_NewDocument(t *testing.T) {
	type input_output struct {
		id          string
		partnerID   string
		title       string
		filePath    string
		extension   string
		lastUpdated time.Time
		createdAt   time.Time
	}

	type testCase struct {
		test           string
		input          input_output
		expectedOutput input_output
		expectedError  error
	}

	testId := uuid.New().String()
	timeNow := time.Now()
	timeBefore := timeNow.Add(time.Minute * -1)

	testsTable := []testCase{
		{
			test:           "Empty PartnerID, Title and file extension error validation",
			input:          input_output{},
			expectedOutput: input_output{},
			expectedError:  errors.New("invalid fields: PartnerDocument.PartnerID: \"\", PartnerDocument.Title: \"\", PartnerDocument.File.Extension: \"\""),
		},
		{
			test: "PartnerDocument ID, PartnerID and Title length error validation",
			input: input_output{
				id:          "Invalid ID",
				partnerID:   "Invalid Partner ID",
				title:       "AA",
				filePath:    "path/to/file.pdf",
				extension:   "pdf",
				createdAt:   timeNow,
				lastUpdated: timeNow,
			},
			expectedOutput: input_output{
				id:          "Invalid ID",
				partnerID:   "Invalid Partner ID",
				title:       "AA",
				filePath:    "path/to/file.pdf",
				extension:   "pdf",
				createdAt:   timeNow,
				lastUpdated: timeNow,
			},
			expectedError: errors.New("invalid fields: PartnerDocument.ID: \"Invalid ID\", PartnerDocument.PartnerID: \"Invalid Partner ID\", PartnerDocument.Title: \"AA\""),
		},
		{
			test: "PartnerDocument CreatedAt and LastUpdated error validation",
			input: input_output{
				id:          testId,
				partnerID:   testId,
				title:       "Partner document",
				filePath:    "path/to/file.pdf",
				extension:   "pdf",
				createdAt:   timeNow,
				lastUpdated: timeBefore,
			},
			expectedOutput: input_output{
				id:          testId,
				partnerID:   testId,
				title:       "Partner document",
				filePath:    "path/to/file.pdf",
				extension:   "pdf",
				createdAt:   timeNow,
				lastUpdated: timeBefore,
			},
			expectedError: fmt.Errorf("invalid fields: PartnerDocument.LastUpdated: \"%s\", PartnerDocument.CreatedAt: \"%s\"", timeBefore, timeNow),
		},
		{
			test: "Valid PartnerDocument fields generating new ID, FilePath, CreatedAt and LastUpdated when empty",
			input: input_output{
				id:          "",
				partnerID:   testId,
				title:       "Partner document",
				filePath:    "",
				extension:   "pdf",
				createdAt:   time.Time{},
				lastUpdated: time.Time{},
			},
			expectedOutput: input_output{
				id:          "", //Must have a new generated UUID
				partnerID:   testId,
				title:       "Partner document",
				filePath:    "", //Must generate a new filePath
				extension:   "pdf",
				createdAt:   time.Time{}, //Must have a CreatedAt with time.Now
				lastUpdated: time.Time{}, //Must have a LastUpdated with time.Now
			},
			expectedError: nil,
		},
		{
			test: "Valid PartnerDocument fields",
			input: input_output{
				id:          testId,
				partnerID:   testId,
				title:       "Partner document",
				filePath:    "path/to/file.pdf",
				extension:   "pdf",
				createdAt:   timeBefore,
				lastUpdated: timeNow,
			},
			expectedOutput: input_output{
				id:          testId,
				partnerID:   testId,
				title:       "Partner document",
				filePath:    "path/to/file.pdf",
				extension:   "pdf",
				createdAt:   timeBefore,
				lastUpdated: timeNow,
			},
			expectedError: nil,
		},
	}

	for _, tc := range testsTable {
		fmt.Printf("Test case: %s\n\n", tc.test)
		compDoc, err := entity.NewDocument(
			tc.input.id,
			tc.input.partnerID,
			tc.input.title,
			tc.input.filePath,
			tc.input.extension,
			tc.input.lastUpdated,
			tc.input.createdAt,
		)

		require.NotEmpty(t, compDoc.ID)

		if tc.input.id != "" {
			require.Equal(t, tc.expectedOutput.id, compDoc.ID)
		}

		require.Equal(t, tc.expectedOutput.partnerID, compDoc.PartnerID)
		require.Equal(t, tc.expectedOutput.title, compDoc.Title)

		require.NotEmpty(t, compDoc.File.FilePath)
		require.Equal(t, tc.expectedOutput.extension, compDoc.File.Extension)

		require.NotZero(t, compDoc.CreatedAt)

		if !tc.input.createdAt.IsZero() {
			require.Equal(t, tc.expectedOutput.createdAt, compDoc.CreatedAt)
		}

		require.NotZero(t, compDoc.LastUpdated)

		if !tc.input.createdAt.IsZero() {
			require.Equal(t, tc.expectedOutput.lastUpdated, compDoc.LastUpdated)
		}

		if tc.expectedError != nil {
			require.Equal(t, tc.expectedError, err)
			continue
		}

		require.Nil(t, err)
	}
}
