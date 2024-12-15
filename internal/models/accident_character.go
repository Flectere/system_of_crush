package models

type AccidentCharacter struct {
	ID             *int           `json:"id,omitempty"`
	Name           *string        `json:"name,omitempty"`
	Specialization Specialization `json:"specialization,omitempty"`
}
