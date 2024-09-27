package models

import (
	"time"
)

type Appeal struct {
	ID             int            `json:"id,omitempty"`
	Title          string         `json:"title"`
	Description    string         `json:"description,omitempty" `
	CreateDate     time.Time      `json:"create_date,omitempty"`
	Specialization Specialization `json:"specialization"`
	Importance     Importance     `json:"importance"`
}
