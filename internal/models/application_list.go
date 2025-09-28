package models

import (
	"time"
)

type ApplicationList struct {
	ID             *int       `json:"id,omitempty"`
	CreateDate     *time.Time `json:"create_date,omitempty"`
	Accident       *string    `json:"accident,omitempty"`
	Importance     *string    `json:"importance,omitempty"`
	Status         *string    `json:"status,omitempty"`
	Address        *string    `json:"address,omitempty"`
	Specialization *string    `json:"specialization,omitempty"`
}
