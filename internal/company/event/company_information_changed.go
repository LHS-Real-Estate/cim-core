package event

import "time"

const CompanyInformationChangedEventName = "CompanyInformationChanged"

type CompanyInformationChanged struct {
	name       string
	payload    interface{}
	occurredAt time.Time
}

func NewCompanyInformationChanged() *CompanyInformationChanged {
	return &CompanyInformationChanged{
		name: CompanyInformationChangedEventName,
	}
}

func (e *CompanyInformationChanged) Name() string {
	return e.name
}

func (e *CompanyInformationChanged) Payload() interface{} {
	return e.payload
}

func (e *CompanyInformationChanged) SetPayload(payload interface{}) {
	e.payload = payload
}

func (e *CompanyInformationChanged) OccurredAt() time.Time {
	return e.occurredAt
}

func (e *CompanyInformationChanged) SetOccurredAt(t time.Time) {
	e.occurredAt = t
}
