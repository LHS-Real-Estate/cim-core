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

func TestPartner_NewPartner(t *testing.T) {
	type input_output struct {
		id        string
		companyID string
		name      string
		surname   string
		isActive  bool
		createdAt time.Time
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
			test:           "Empty CompanyID and Name error validation",
			input:          input_output{id: testId, createdAt: timeNow},
			expectedOutput: input_output{id: testId, createdAt: timeNow},
			expectedError:  errors.New("invalid fields: Partner.CompanyID: \"\", Partner.Name: \"\""),
		},
		{
			test: "Company Name and Surname length error validation",
			input: input_output{
				id:        testId,
				companyID: testId,
				name:      "A",
				surname:   "AA",
				isActive:  true,
				createdAt: timeNow,
			},
			expectedOutput: input_output{
				id:        testId,
				companyID: testId,
				name:      "A",
				surname:   "AA",
				isActive:  true,
				createdAt: timeNow,
			},
			expectedError: errors.New("invalid fields: Partner.Name: \"A\", Partner.Surname: \"AA\""),
		},
		{
			test: "Partner ID and Company ID error validation",
			input: input_output{
				id:        "Invalid UUID",
				companyID: "Invalid UUID",
				name:      "John",
				surname:   "Doe",
				isActive:  true,
				createdAt: timeNow,
			},
			expectedOutput: input_output{
				id:        "Invalid UUID",
				companyID: "Invalid UUID",
				name:      "John",
				surname:   "Doe",
				isActive:  true,
				createdAt: timeNow,
			},
			expectedError: errors.New("invalid fields: Partner.ID: \"Invalid UUID\", Partner.CompanyID: \"Invalid UUID\""),
		},
		{
			test: "Valid Company fields generating new ID and CreatedAt when empty",
			input: input_output{
				id:        "",
				companyID: testId,
				name:      "John",
				surname:   "Doe",
				isActive:  true,
				createdAt: time.Time{},
			},
			expectedOutput: input_output{
				id:        "", //Must have a new generated UUID
				companyID: testId,
				name:      "John",
				surname:   "Doe",
				isActive:  true,
				createdAt: time.Time{}, //Must have a CreatedAt with time.Now
			},
			expectedError: nil,
		},
		{
			test: "Valid Company fields",
			input: input_output{
				companyID: testId,
				name:      "John",
				surname:   "Doe",
				isActive:  true,
				createdAt: timeNow,
			},
			expectedOutput: input_output{
				companyID: testId,
				name:      "John",
				surname:   "Doe",
				isActive:  true,
				createdAt: timeNow,
			},
			expectedError: nil,
		},
	}

	for _, tc := range testsTable {
		fmt.Printf("Test case: %s\n\n", tc.test)
		partner, err := entity.NewPartner(
			tc.input.id,
			tc.input.companyID,
			tc.input.name,
			tc.input.surname,
			tc.input.isActive,
			tc.input.createdAt,
		)

		require.NotEmpty(t, partner.ID)

		if tc.input.id != "" {
			require.Equal(t, tc.expectedOutput.id, partner.ID)
		}

		require.Equal(t, tc.expectedOutput.companyID, partner.CompanyID)
		require.Equal(t, tc.expectedOutput.name, partner.Name)
		require.Equal(t, tc.expectedOutput.surname, partner.Surname)
		require.Equal(t, tc.expectedOutput.isActive, partner.IsActive)

		require.NotZero(t, partner.CreatedAt)

		if !tc.input.createdAt.IsZero() {
			require.Equal(t, tc.expectedOutput.createdAt, partner.CreatedAt)
		}

		if tc.expectedError != nil {
			require.Equal(t, tc.expectedError, err)
			continue
		}

		require.Nil(t, err)
	}
}
