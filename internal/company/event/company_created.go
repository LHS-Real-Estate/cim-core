package event

import "time"

const CompanyCreatedEventName = "CompanyCreated"

type CompanyCreated struct {
	name       string
	payload    interface{}
	occurredAt time.Time
}

func NewCompanyCreated() *CompanyCreated {
	return &CompanyCreated{
		name: CompanyCreatedEventName,
	}
}

func (e *CompanyCreated) Name() string {
	return e.name
}

func (e *CompanyCreated) Payload() interface{} {
	return e.payload
}

func (e *CompanyCreated) SetPayload(payload interface{}) {
	e.payload = payload
}

func (e *CompanyCreated) OccurredAt() time.Time {
	return e.occurredAt
}

func (e *CompanyCreated) SetOccurredAt(t time.Time) {
	e.occurredAt = t
}
