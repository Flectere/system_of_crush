package models

import (
	"time"
)

type ApplicationBrigadir struct {
	ID              *int       `json:"id"`
	Address         *string    `json:"address"`
	Accident        *string    `json:"accident"`
	Importance      *string    `json:"importance"`
	StartDate       *time.Time `json:"start_date"`
	Character       *string    `json:"character"`
	Material        *string    `json:"material"`
	DamageType      *string    `json:"damage_type"`
	DamagePoint     *string    `json:"damage_point"`
	Description     *string    `json:"description"`
	RecommendedTime *int       `json:"recommended_time"`
}
