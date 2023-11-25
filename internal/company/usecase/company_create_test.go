package usecase_test

import (
	"errors"
	"time"

	company_event "github.com/LHS-Real-Estate/cim-core/internal/company/event"
	"github.com/LHS-Real-Estate/cim-core/internal/company/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CompanyCreateUseCaseTestSuite struct {
	suite.Suite
	MockCompanyRepository *MockCompanyRepository
	CompanyCreatedEvent   *company_event.CompanyCreated
	MockDispatcher        *MockEventDispatcher
	UseCase               *usecase.CompanyCreateUseCase
}

func (suite *CompanyCreateUseCaseTestSuite) SetupTest() {
	suite.MockCompanyRepository = &MockCompanyRepository{}
	suite.CompanyCreatedEvent = company_event.NewCompanyCreated()
	suite.MockDispatcher = &MockEventDispatcher{}
	suite.UseCase = usecase.NewCompanyCreateUseCase(
		suite.MockCompanyRepository,
		*suite.CompanyCreatedEvent,
		suite.MockDispatcher,
	)
}

func (suite *CompanyCreateUseCaseTestSuite) TestCompanyCreateUseCase_Execute_CompanyEntityError() {
	//Invalid Company Input
	input := usecase.CompanyCreateInputDTO{
		EIN:                   "A",
		Name:                  "A",
		FullName:              "A",
		MunicipalRegistration: "",
		StateRegistration:     "",
	}

	output, err := suite.UseCase.Execute(input)
	suite.NotNil(err)
	suite.Empty(output)
	suite.MockCompanyRepository.AssertNumberOfCalls(suite.T(), "Create", 0)

	suite.Empty(suite.UseCase.CompanyCreatedEvent.Payload())
	suite.True(suite.UseCase.CompanyCreatedEvent.OccurredAt().IsZero())
	suite.MockDispatcher.AssertNumberOfCalls(suite.T(), "Dispatch", 0)
}

func (suite *CompanyCreateUseCaseTestSuite) TestCompanyCreateUseCase_Execute_CompanyRepositoryCreateError() {
	input := usecase.CompanyCreateInputDTO{
		EIN:                   "12.345.678/0001-90",
		Name:                  "Company Test",
		FullName:              "The Company Test Inc.",
		MunicipalRegistration: "1.234.567/001-8",
		StateRegistration:     "123456789.00-12",
	}

	suite.MockCompanyRepository.On("Create", mock.AnythingOfType("*entity.Company")).Return(
		errors.New("a CompanyRepository create error"),
	)
	suite.MockDispatcher.On("Dispatch", mock.AnythingOfType("event.EventInterface"))

	output, err := suite.UseCase.Execute(input)
	suite.NotNil(err)
	suite.Empty(output)
	suite.MockCompanyRepository.AssertNumberOfCalls(suite.T(), "Create", 1)

	suite.Empty(suite.UseCase.CompanyCreatedEvent.Payload())
	suite.True(suite.UseCase.CompanyCreatedEvent.OccurredAt().IsZero())
	suite.MockDispatcher.AssertNumberOfCalls(suite.T(), "Dispatch", 0)
}

func (suite *CompanyCreateUseCaseTestSuite) TestCompanyCreateUseCase_Execute() {
	validInput := usecase.CompanyCreateInputDTO{
		EIN:                   "12.345.678/0001-90",
		Name:                  "Company Test",
		FullName:              "The Company Test Inc.",
		MunicipalRegistration: "1.234.567/001-8",
		StateRegistration:     "123456789.00-12",
	}

	expectedOutput := usecase.CompanyCreateOutputDTO{
		ID:                    "VALID UUID",
		EIN:                   validInput.EIN,
		Name:                  validInput.Name,
		FullName:              validInput.FullName,
		MunicipalRegistration: validInput.MunicipalRegistration,
		StateRegistration:     validInput.StateRegistration,
		CreatedAt:             time.Now(),
	}

	suite.MockCompanyRepository.On("Create", mock.AnythingOfType("*entity.Company")).Return(nil)
	suite.MockDispatcher.On("Dispatch", mock.Anything)

	output, err := suite.UseCase.Execute(validInput)
	suite.Nil(err)
	suite.NotNil(output.ID)
	suite.Equal(expectedOutput.EIN, output.EIN)
	suite.Equal(expectedOutput.Name, output.Name)
	suite.Equal(expectedOutput.FullName, output.FullName)
	suite.Equal(expectedOutput.MunicipalRegistration, output.MunicipalRegistration)
	suite.Equal(expectedOutput.StateRegistration, output.StateRegistration)
	suite.GreaterOrEqual(output.CreatedAt, expectedOutput.CreatedAt)
	suite.MockCompanyRepository.AssertNumberOfCalls(suite.T(), "Create", 1)

	suite.Equal(output, suite.UseCase.CompanyCreatedEvent.Payload()) //Payload should be equal output DTO
	suite.GreaterOrEqual(suite.UseCase.CompanyCreatedEvent.OccurredAt(), output.CreatedAt)
	suite.MockDispatcher.AssertNumberOfCalls(suite.T(), "Dispatch", 1)
}
