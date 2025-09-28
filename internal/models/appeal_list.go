package models

import (
	"time"
)

type AppealList struct {
	ID             *int       `json:"id,omitempty"`
	CreateDate     *time.Time `json:"create_date,omitempty"`
	Accident       *string    `json:"accident"`
	Address        *string    `json:"address"`
	Status         *bool      `json:"status,omitempty"`
	Specialization *string    `json:"specialization"`
}
