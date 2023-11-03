package entity_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/LHS-Real-Estate/cim-core/internal/company/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCompany_NewCompany(t *testing.T) {
	type input_output struct {
		id                    string
		ein                   string
		name                  string
		fullName              string
		municipalRegistration string
		stateRegistration     string
		createdAt             time.Time
	}

	type testCase struct {
		test           string
		input          input_output
		expectedOutput input_output
		expectedError  error
	}

	testId := uuid.New().String()
	timeNow := time.Now()

	testsTable := []testCase{
		{
			test:           "Empty EIN, Name and FullName validation",
			input:          input_output{id: testId, createdAt: timeNow},
			expectedOutput: input_output{id: testId, createdAt: timeNow},
			expectedError:  errors.New("invalid fields: Company.EIN: \"\", Company.Name: \"\", Company.FullName: \"\""),
		},
		{
			test: "Company Name and FullName length validation",
			input: input_output{
				id:                    testId,
				ein:                   "01.234.567/0001-89",
				name:                  "AA",
				fullName:              "AA",
				municipalRegistration: "0123456/001-7",
				stateRegistration:     "012345678.90-12",
				createdAt:             timeNow,
			},
			expectedOutput: input_output{
				id:                    testId,
				ein:                   "01.234.567/0001-89",
				name:                  "AA",
				fullName:              "AA",
				municipalRegistration: "0123456/001-7",
				stateRegistration:     "012345678.90-12",
				createdAt:             timeNow,
			},
			expectedError: errors.New("invalid fields: Company.Name: \"\", Company.FullName: \"\""),
		},
		{
			test: "Company ID validation",
			input: input_output{
				id:                    "Invalid UUID",
				ein:                   "01.234.567/0001-89",
				name:                  "Company Test",
				fullName:              "Company Test Inc",
				municipalRegistration: "0123456/001-7",
				stateRegistration:     "012345678.90-12",
				createdAt:             timeNow,
			},
			expectedOutput: input_output{
				id:                    "Invalid UUID",
				ein:                   "01.234.567/0001-89",
				name:                  "Company Test",
				fullName:              "Company Test Inc",
				municipalRegistration: "0123456/001-7",
				stateRegistration:     "012345678.90-12",
				createdAt:             timeNow,
			},
			expectedError: errors.New("invalid fields: Company.ID: \"\""),
		},
		{
			test: "Valid Company fields generating new ID and CreatedAt when empty",
			input: input_output{
				id:                    "",
				ein:                   "01.234.567/0001-89",
				name:                  "Company Test",
				fullName:              "Company Test Inc",
				municipalRegistration: "0123456/001-7",
				stateRegistration:     "012345678.90-12",
				createdAt:             time.Time{},
			},
			expectedOutput: input_output{
				id:                    "", //Must have a new generated UUID
				ein:                   "01.234.567/0001-89",
				name:                  "Company Test",
				fullName:              "Company Test Inc",
				municipalRegistration: "0123456/001-7",
				stateRegistration:     "012345678.90-12",
				createdAt:             time.Time{}, //Must have a CreatedAt with time.Now
			},
			expectedError: nil,
		},
		{
			test: "Valid Company fields",
			input: input_output{
				id:                    testId,
				ein:                   "01.234.567/0001-89",
				name:                  "Company Test",
				fullName:              "Company Test Inc",
				municipalRegistration: "0123456/001-7",
				stateRegistration:     "012345678.90-12",
				createdAt:             timeNow,
			},
			expectedOutput: input_output{
				id:                    testId,
				ein:                   "01.234.567/0001-89",
				name:                  "Company Test",
				fullName:              "Company Test Inc",
				municipalRegistration: "0123456/001-7",
				stateRegistration:     "012345678.90-12",
				createdAt:             timeNow,
			},
			expectedError: nil,
		},
	}

	for _, tc := range testsTable {
		fmt.Printf("Test case: %s", tc.test)
		comp, err := entity.NewCompany(
			tc.input.id,
			tc.input.ein,
			tc.input.name,
			tc.input.fullName,
			tc.input.municipalRegistration,
			tc.input.stateRegistration,
			tc.input.createdAt,
		)

		assert.NotEmpty(t, comp.ID)

		if tc.input.id != "" {
			assert.Equal(t, comp.ID, tc.expectedOutput.id)
		}

		assert.Equal(t, comp.EIN, tc.expectedOutput.ein)
		assert.Equal(t, comp.Name, tc.expectedOutput.name)
		assert.Equal(t, comp.FullName, tc.expectedOutput.fullName)
		assert.Equal(t, comp.MunicipalRegistration, tc.expectedOutput.municipalRegistration)
		assert.Equal(t, comp.StateRegistration, tc.expectedOutput.stateRegistration)

		assert.NotZero(t, comp.CreatedAt)

		if !tc.input.createdAt.IsZero() {
			assert.Equal(t, comp.CreatedAt, tc.expectedOutput.createdAt)
		}

		if tc.expectedError != nil {
			assert.Error(t, err, tc.expectedError)
			continue
		}

		assert.Nil(t, err)
	}
}
