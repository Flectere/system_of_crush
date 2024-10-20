package models

import (
	"time"
)

type Application struct {
	ID          int             `json:"id,omitempty"`
	Description string          `json:"description,omitempty" `
	CreateDate  time.Time       `json:"create_date,omitempty"`
	Accident    AccidentContent `json:"accident"`
	Importance  Importance      `json:"importance"`
	Status      Status          `json:"status"`
}
