package usecase

import (
	"time"

	"github.com/LHS-Real-Estate/cim-core/internal/company/entity"
	company_event "github.com/LHS-Real-Estate/cim-core/internal/company/event"
	ed "github.com/yamauthi/event-dispatcher"
)

type CreateCompanyInputDTO struct {
	EIN                   string
	Name                  string
	FullName              string
	MunicipalRegistration string
	StateRegistration     string
}

type CreateCompanyOutputDTO struct {
	ID                    string
	EIN                   string
	Name                  string
	FullName              string
	MunicipalRegistration string
	StateRegistration     string
	CreatedAt             time.Time
}

type CreateCompanyUseCase struct {
	CompanyRepository   entity.CompanyRepositoryInterface
	CompanyCreatedEvent company_event.CompanyCreated
	Dispatcher          ed.EventDispatcherInterface
}

func NewCreateCompanyUseCase(
	CompanyRepository entity.CompanyRepositoryInterface,
	CompanyCreatedEvent company_event.CompanyCreated,
	Dispatcher ed.EventDispatcherInterface,
) *CreateCompanyUseCase {
	return &CreateCompanyUseCase{
		CompanyRepository:   CompanyRepository,
		CompanyCreatedEvent: CompanyCreatedEvent,
		Dispatcher:          Dispatcher,
	}
}

func (uc *CreateCompanyUseCase) Execute(input CreateCompanyInputDTO) (CreateCompanyOutputDTO, error) {
	c, err := entity.NewCompany(
		"", // Will generate ID
		input.EIN,
		input.Name,
		input.FullName,
		input.MunicipalRegistration,
		input.StateRegistration,
		time.Time{}, //Will fill CreatedAt
	)

	if err != nil {
		return CreateCompanyOutputDTO{}, err
	}

	err = uc.CompanyRepository.Create(&c)
	if err != nil {
		return CreateCompanyOutputDTO{}, err
	}

	dto := CreateCompanyOutputDTO{
		ID:                    c.ID,
		EIN:                   c.EIN,
		Name:                  c.Name,
		FullName:              c.FullName,
		MunicipalRegistration: c.MunicipalRegistration,
		StateRegistration:     c.StateRegistration,
		CreatedAt:             c.CreatedAt,
	}

	uc.CompanyCreatedEvent.SetPayload(dto)
	uc.CompanyCreatedEvent.SetOccurredAt(time.Now())
	uc.Dispatcher.Dispatch(&uc.CompanyCreatedEvent)

	return dto, nil
}
