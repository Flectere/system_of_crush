package models

import (
	"time"
)

type Application struct {
	ID            *int           `json:"id,omitempty"`
	Description   *string        `json:"description,omitempty" `
	CreateDate    *time.Time     `json:"create_date,omitempty"`
	Accident      Accident       `json:"accident,omitempty"`
	Importance    Importance     `json:"importance,omitempty"`
	Status        Status         `json:"status,omitempty"`
	Address       *string        `json:"address,omitempty"`
	AccidentCause *string        `json:"cause,omitempty"`
	Material      Material       `json:"material,omitempty"`
	DamageType    Damage         `json:"damage_type,omitempty"`
	DamagePoint   *string        `json:"damage_point,omitempty"`
	Brigade       Brigade        `json:"brigade,omitempty"`
	Duration      *time.Duration `json:"duration,omitempty"`
	Image         []byte         `json:"image,omitempty"`
}
