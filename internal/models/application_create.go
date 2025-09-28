package models

import (
	"time"
)

type ApplicationCreate struct {
	Description   *string    `json:"description,omitempty" `
	CreateDate    *time.Time `json:"create_date,omitempty"`
	Accident      *int       `json:"id_accident,omitempty"`
	Importance    *int       `json:"id_importance,omitempty"`
	Status        *int       `json:"id_status,omitempty"`
	Address       *string    `json:"address,omitempty"`
	AccidentCause *string    `json:"cause,omitempty"`
	Material      *int       `json:"id_material,omitempty"`
	DamageType    *int       `json:"id_damage_type,omitempty"`
	DamagePoint   *string    `json:"damage_point,omitempty"`
	Operator      *int       `json:"id_operator,omitempty"`
}
