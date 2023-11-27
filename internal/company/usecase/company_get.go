package usecase

import (
	"errors"
	"reflect"
	"time"

	"github.com/LHS-Real-Estate/cim-core/internal/company/entity"
)

var ErrCompanyGetUseCaseEmptyID = errors.New("field Company.ID must not be empty")
var ErrCompanyGetUseCaseNotFound = errors.New("no Company found for given Company.ID")

type CompanyGetInputDTO struct {
	ID string
}

type CompanyGetOutputDTO struct {
	ID                    string
	EIN                   string
	Name                  string
	FullName              string
	MunicipalRegistration string
	StateRegistration     string
	CreatedAt             time.Time
}

type CompanyGetUseCase struct {
	CompanyRepository entity.CompanyRepositoryInterface
}

func NewCompanyGetUseCase(CompanyRepository entity.CompanyRepositoryInterface) *CompanyGetUseCase {
	return &CompanyGetUseCase{CompanyRepository: CompanyRepository}
}

func (uc *CompanyGetUseCase) Execute(input CompanyGetInputDTO) (CompanyGetOutputDTO, error) {
	if input.ID == "" {
		return CompanyGetOutputDTO{}, ErrCompanyGetUseCaseEmptyID
	}

	c, err := uc.CompanyRepository.Get(input.ID)
	if err != nil {
		return CompanyGetOutputDTO{}, err
	}

	if reflect.DeepEqual(entity.Company{}, c) {
		return CompanyGetOutputDTO{}, ErrCompanyGetUseCaseNotFound
	}

	dto := CompanyGetOutputDTO{
		ID:                    c.ID,
		EIN:                   c.EIN,
		Name:                  c.Name,
		FullName:              c.FullName,
		MunicipalRegistration: c.MunicipalRegistration,
		StateRegistration:     c.StateRegistration,
		CreatedAt:             c.CreatedAt,
	}

	return dto, nil
}
