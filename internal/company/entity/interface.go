package entity

type CompanyRepositoryInterface interface {
	ChangeInformation(company *Company) error
	Create(company *Company) error
	Get(companyID string) (Company, error)
}
