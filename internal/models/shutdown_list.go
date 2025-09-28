package models

import "time"

type ShutdownList struct {
	ID             *int       `json:"id"`
	Address        *string    `json:"address"`
	Accident       *string    `json:"accident"`
	Date           *time.Time `json:"date"`
	DayCount       *int       `json:"day_count"`
	Specialization *string    `json:"specialization"`
	Type           *string    `json:"type"`
}
