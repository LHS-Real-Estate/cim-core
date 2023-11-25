package usecase

import (
	"time"

	"github.com/LHS-Real-Estate/cim-core/internal/company/entity"
	company_event "github.com/LHS-Real-Estate/cim-core/internal/company/event"
	ed "github.com/yamauthi/event-dispatcher"
)

type CompanyChangeInformationInputDTO struct {
	ID                    string
	EIN                   string
	Name                  string
	FullName              string
	MunicipalRegistration string
	StateRegistration     string
}

type CompanyChangeInformationOutputDTO struct {
	ID                    string
	EIN                   string
	Name                  string
	FullName              string
	MunicipalRegistration string
	StateRegistration     string
}

type CompanyChangeInformationUseCase struct {
	CompanyRepository              entity.CompanyRepositoryInterface
	CompanyInformationChangedEvent company_event.CompanyInformationChanged
	Dispatcher                     ed.EventDispatcherInterface
}

func NewCompanyChangeInformationUseCase(
	CompanyRepository entity.CompanyRepositoryInterface,
	CompanyInformationChangedEvent company_event.CompanyInformationChanged,
	Dispatcher ed.EventDispatcherInterface,
) *CompanyChangeInformationUseCase {
	return &CompanyChangeInformationUseCase{
		CompanyRepository:              CompanyRepository,
		CompanyInformationChangedEvent: CompanyInformationChangedEvent,
		Dispatcher:                     Dispatcher,
	}
}

func (uc *CompanyChangeInformationUseCase) Execute(input CompanyChangeInformationInputDTO) (CompanyChangeInformationOutputDTO, error) {
	c, err := entity.NewCompany(
		input.ID,
		input.EIN,
		input.Name,
		input.FullName,
		input.MunicipalRegistration,
		input.StateRegistration,
		time.Time{}, //Ignored for this use case
	)

	if err != nil {
		return CompanyChangeInformationOutputDTO{}, err
	}

	err = uc.CompanyRepository.ChangeInformation(&c)
	if err != nil {
		return CompanyChangeInformationOutputDTO{}, err
	}

	dto := CompanyChangeInformationOutputDTO{
		ID:                    c.ID,
		EIN:                   c.EIN,
		Name:                  c.Name,
		FullName:              c.FullName,
		MunicipalRegistration: c.MunicipalRegistration,
		StateRegistration:     c.StateRegistration,
	}

	uc.CompanyInformationChangedEvent.SetPayload(dto)
	uc.CompanyInformationChangedEvent.SetOccurredAt(time.Now())
	uc.Dispatcher.Dispatch(&uc.CompanyInformationChangedEvent)

	return dto, nil
}
