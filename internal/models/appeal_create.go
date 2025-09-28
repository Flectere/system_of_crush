package models

import (
	"time"
)

type AppealCreate struct {
	Description      *string    `json:"description,omitempty" `
	CreateDate       *time.Time `json:"create_date,omitempty"`
	Accident         *int       `json:"id_accident"`
	Address          *string    `json:"address"`
	ApplicantName    *string    `json:"applicant_name,omitempty"`
	ApplicantNumber  *string    `json:"applicant_number,omitempty"`
	AdditionalNumber *string    `json:"additional_number,omitempty"`
	Application      *int       `json:"id_application,omitempty"`
}
