package usecase_test

import (
	"errors"
	"time"

	company_event "github.com/LHS-Real-Estate/cim-core/internal/company/event"
	"github.com/LHS-Real-Estate/cim-core/internal/company/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CompanyChangeInformationUseCaseTestSuite struct {
	suite.Suite
	MockCompanyRepository          *MockCompanyRepository
	CompanyInformationChangedEvent *company_event.CompanyInformationChanged
	MockDispatcher                 *MockEventDispatcher
	UseCase                        *usecase.CompanyChangeInformationUseCase
}

func (suite *CompanyChangeInformationUseCaseTestSuite) SetupTest() {
	suite.MockCompanyRepository = &MockCompanyRepository{}
	suite.CompanyInformationChangedEvent = company_event.NewCompanyInformationChanged()
	suite.MockDispatcher = &MockEventDispatcher{}
	suite.UseCase = usecase.NewCompanyChangeInformationUseCase(
		suite.MockCompanyRepository,
		*suite.CompanyInformationChangedEvent,
		suite.MockDispatcher,
	)
}

func (suite *CompanyChangeInformationUseCaseTestSuite) TestCompanyChangeInformationUseCase_Execute_CompanyEntityError() {
	//Invalid Company Input
	input := usecase.CompanyChangeInformationInputDTO{
		ID:                    "Invalid UUID",
		EIN:                   "A",
		Name:                  "A",
		FullName:              "A",
		MunicipalRegistration: "",
		StateRegistration:     "",
	}

	output, err := suite.UseCase.Execute(input)
	suite.NotNil(err)
	suite.Empty(output)
	suite.MockCompanyRepository.AssertNumberOfCalls(suite.T(), "ChangeInformation", 0)

	suite.Empty(suite.UseCase.CompanyInformationChangedEvent.Payload())
	suite.True(suite.UseCase.CompanyInformationChangedEvent.OccurredAt().IsZero())
	suite.MockDispatcher.AssertNumberOfCalls(suite.T(), "Dispatch", 0)
}

func (suite *CompanyChangeInformationUseCaseTestSuite) TestCompanyChangeInformationUseCase_Execute_CompanyRepositoryChangeInformationError() {
	input := usecase.CompanyChangeInformationInputDTO{
		ID:                    "914e6994-9dda-4c17-810d-d0fbdb57f4dc",
		EIN:                   "12.345.678/0001-90",
		Name:                  "Company Test",
		FullName:              "The Company Test Inc.",
		MunicipalRegistration: "1.234.567/001-8",
		StateRegistration:     "123456789.00-12",
	}

	suite.MockCompanyRepository.On("ChangeInformation", mock.AnythingOfType("*entity.Company")).Return(
		errors.New("a CompanyRepository changeInformation error"),
	)
	suite.MockDispatcher.On("Dispatch", mock.Anything)

	output, err := suite.UseCase.Execute(input)
	suite.NotNil(err)
	suite.Empty(output)
	suite.MockCompanyRepository.AssertNumberOfCalls(suite.T(), "ChangeInformation", 1)

	suite.Empty(suite.UseCase.CompanyInformationChangedEvent.Payload())
	suite.True(suite.UseCase.CompanyInformationChangedEvent.OccurredAt().IsZero())
	suite.MockDispatcher.AssertNumberOfCalls(suite.T(), "Dispatch", 0)
}

func (suite *CompanyChangeInformationUseCaseTestSuite) TestCompanyChangeInformationUseCase_Execute() {
	validInput := usecase.CompanyChangeInformationInputDTO{
		ID:                    "914e6994-9dda-4c17-810d-d0fbdb57f4dc",
		EIN:                   "12.345.678/0001-90",
		Name:                  "Company Test",
		FullName:              "The Company Test Inc.",
		MunicipalRegistration: "1.234.567/001-8",
		StateRegistration:     "123456789.00-12",
	}

	expectedOutput := usecase.CompanyChangeInformationOutputDTO{
		ID:                    "914e6994-9dda-4c17-810d-d0fbdb57f4dc",
		EIN:                   "12.345.678/0001-90",
		Name:                  "Company Test",
		FullName:              "The Company Test Inc.",
		MunicipalRegistration: "1.234.567/001-8",
		StateRegistration:     "123456789.00-12",
	}

	timeBeforeChangeInformation := time.Now()

	suite.MockCompanyRepository.On("ChangeInformation", mock.AnythingOfType("*entity.Company")).Return(
		nil,
	)
	suite.MockDispatcher.On("Dispatch", mock.Anything)

	output, err := suite.UseCase.Execute(validInput)
	suite.Nil(err)
	suite.NotNil(output.ID)
	suite.Equal(expectedOutput.EIN, output.EIN)
	suite.Equal(expectedOutput.Name, output.Name)
	suite.Equal(expectedOutput.FullName, output.FullName)
	suite.Equal(expectedOutput.MunicipalRegistration, output.MunicipalRegistration)
	suite.Equal(expectedOutput.StateRegistration, output.StateRegistration)
	suite.MockCompanyRepository.AssertNumberOfCalls(suite.T(), "ChangeInformation", 1)

	suite.Equal(output, suite.UseCase.CompanyInformationChangedEvent.Payload()) //Payload should be equal output DTO
	suite.False(suite.UseCase.CompanyInformationChangedEvent.OccurredAt().IsZero())
	suite.GreaterOrEqual(suite.UseCase.CompanyInformationChangedEvent.OccurredAt(), timeBeforeChangeInformation)
	suite.MockDispatcher.AssertNumberOfCalls(suite.T(), "Dispatch", 1)
}
