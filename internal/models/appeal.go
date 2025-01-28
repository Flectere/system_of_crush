package models

import (
	"time"
)

type Appeal struct {
	ID               *int        `json:"id,omitempty"`
	Description      *string     `json:"description,omitempty" `
	CreateDate       *time.Time  `json:"create_date,omitempty"`
	Accident         Accident    `json:"accident"`
	Address          *string     `json:"address"`
	ApplicantName    *string     `json:"applicant_name,omitempty"`
	ApplicantNumber  *string     `json:"applicant_number,omitempty"`
	AdditionalNumber *string     `json:"additional_number,omitempty"`
	Application      Application `json:"application,omitempty"`
	IsActive         *bool       `json:"is_active,omitempty"`
}
