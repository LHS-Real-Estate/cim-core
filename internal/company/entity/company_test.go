package entity_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/LHS-Real-Estate/cim-core/internal/company/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
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
			test:           "Empty EIN, Name and FullName error validation",
			input:          input_output{id: testId, createdAt: timeNow},
			expectedOutput: input_output{id: testId, createdAt: timeNow},
			expectedError:  errors.New("invalid fields: Company.EIN: \"\", Company.Name: \"\", Company.FullName: \"\""),
		},
		{
			test: "Company Name and FullName length error validation",
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
			expectedError: errors.New("invalid fields: Company.Name: \"AA\", Company.FullName: \"AA\""),
		},
		{
			test: "Company ID error validation",
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
			expectedError: errors.New("invalid fields: Company.ID: \"Invalid UUID\""),
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
		fmt.Printf("Test case: %s\n\n", tc.test)
		comp, err := entity.NewCompany(
			tc.input.id,
			tc.input.ein,
			tc.input.name,
			tc.input.fullName,
			tc.input.municipalRegistration,
			tc.input.stateRegistration,
			tc.input.createdAt,
		)

		require.NotEmpty(t, comp.ID)

		if tc.input.id != "" {
			require.Equal(t, tc.expectedOutput.id, comp.ID)
		}

		require.Equal(t, tc.expectedOutput.ein, comp.EIN)
		require.Equal(t, tc.expectedOutput.name, comp.Name)
		require.Equal(t, tc.expectedOutput.fullName, comp.FullName)
		require.Equal(t, tc.expectedOutput.municipalRegistration, comp.MunicipalRegistration)
		require.Equal(t, tc.expectedOutput.stateRegistration, comp.StateRegistration)

		require.NotZero(t, comp.CreatedAt)

		if !tc.input.createdAt.IsZero() {
			require.Equal(t, tc.expectedOutput.createdAt, comp.CreatedAt)
		}

		if tc.expectedError != nil {
			require.Equal(t, tc.expectedError, err)
			continue
		}

		require.Nil(t, err)
	}
}
