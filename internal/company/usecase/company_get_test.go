package usecase_test

import (
	"errors"
	"time"

	"github.com/LHS-Real-Estate/cim-core/internal/company/entity"
	"github.com/LHS-Real-Estate/cim-core/internal/company/usecase"
	"github.com/stretchr/testify/suite"
)

type CompanyGetUseCaseTestSuite struct {
	suite.Suite
	MockCompanyRepository *MockCompanyRepository
	UseCase               *usecase.CompanyGetUseCase
}

func (suite *CompanyGetUseCaseTestSuite) SetupTest() {
	suite.MockCompanyRepository = &MockCompanyRepository{}
	suite.UseCase = usecase.NewCompanyGetUseCase(suite.MockCompanyRepository)
}

func (suite *CompanyGetUseCaseTestSuite) TestCompanyGetUseCase_Execute_EmptyIDError() {
	input := usecase.CompanyGetInputDTO{ID: ""}

	output, err := suite.UseCase.Execute(input)
	suite.Equal(usecase.ErrCompanyGetUseCaseEmptyID, err)
	suite.Empty(output)
	suite.MockCompanyRepository.AssertNumberOfCalls(suite.T(), "Get", 0)
}

func (suite *CompanyGetUseCaseTestSuite) TestCompanyGetUseCase_Execute_CompanyRepositoryGetError() {
	input := usecase.CompanyGetInputDTO{
		ID: "914e6994-9dda-4c17-810d-d0fbdb57f4dc",
	}

	suite.MockCompanyRepository.On("Get", input.ID).Return(
		entity.Company{},
		errors.New("a CompanyRepository Get error"),
	)

	output, err := suite.UseCase.Execute(input)
	suite.NotNil(err)
	suite.Empty(output)
	suite.MockCompanyRepository.AssertNumberOfCalls(suite.T(), "Get", 1)
}

func (suite *CompanyGetUseCaseTestSuite) TestCompanyGetUseCase_Execute_NoCompanyFoundError() {
	input := usecase.CompanyGetInputDTO{
		ID: "Valid but non-existent CompanyID",
	}

	suite.MockCompanyRepository.On("Get", input.ID).Return(
		entity.Company{},
		nil,
	)

	output, err := suite.UseCase.Execute(input)
	suite.Equal(usecase.ErrCompanyGetUseCaseNotFound, err)
	suite.Empty(output)
	suite.MockCompanyRepository.AssertNumberOfCalls(suite.T(), "Get", 1)
}

func (suite *CompanyGetUseCaseTestSuite) TestCompanyGetUseCase_Execute() {
	input := usecase.CompanyGetInputDTO{
		ID: "914e6994-9dda-4c17-810d-d0fbdb57f4dc",
	}

	expectedCompany := entity.Company{
		ID:                    "914e6994-9dda-4c17-810d-d0fbdb57f4dc",
		EIN:                   "12.345.678/0001-90",
		Name:                  "Company Test",
		FullName:              "The Company Test Inc.",
		MunicipalRegistration: "1.234.567/001-8",
		StateRegistration:     "123456789.00-12",
		CreatedAt:             time.Now(),
	}

	suite.MockCompanyRepository.On("Get", input.ID).Return(
		expectedCompany,
		nil,
	)

	output, err := suite.UseCase.Execute(input)
	suite.Nil(err)
	suite.Equal(expectedCompany.ID, output.ID)
	suite.Equal(expectedCompany.EIN, output.EIN)
	suite.Equal(expectedCompany.Name, output.Name)
	suite.Equal(expectedCompany.FullName, output.FullName)
	suite.Equal(expectedCompany.MunicipalRegistration, output.MunicipalRegistration)
	suite.Equal(expectedCompany.StateRegistration, output.StateRegistration)
	suite.Equal(expectedCompany.CreatedAt, output.CreatedAt)
	suite.MockCompanyRepository.AssertNumberOfCalls(suite.T(), "Get", 1)
}
