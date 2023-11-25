package usecase_test

import (
	"testing"

	"github.com/LHS-Real-Estate/cim-core/internal/company/entity"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/yamauthi/event-dispatcher"
)

type MockCompanyRepository struct {
	mock.Mock
}

func (r *MockCompanyRepository) ChangeInformation(company *entity.Company) error {
	args := r.Called(company)
	return args.Error(0)
}
func (r *MockCompanyRepository) Create(company *entity.Company) error {
	args := r.Called(company)
	return args.Error(0)
}
func (r *MockCompanyRepository) Get(companyID string) (entity.Company, error) {
	args := r.Called(companyID)
	return args.Get(0).(entity.Company), args.Error(1)
}

type MockEventDispatcher struct {
	mock.Mock
}

func (ed *MockEventDispatcher) Clear() {}
func (ed *MockEventDispatcher) Dispatch(event event.EventInterface) {
	ed.Called(mock.AnythingOfType("event.EventInterface"))
}
func (ed *MockEventDispatcher) Has(eventName string, handler event.EventHandlerInterface) bool {
	return true
}
func (ed *MockEventDispatcher) Register(eventName string, handler event.EventHandlerInterface) error {
	return nil
}
func (ed *MockEventDispatcher) Remove(eventName string, handler event.EventHandlerInterface) {}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CompanyCreateUseCaseTestSuite))
}
