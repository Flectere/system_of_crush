package models

import "time"

type Shutdown struct {
	ID          *int        `json:"id"`
	Address     *string     `json:"address"`
	Accident    Accident    `json:"accident"`
	Application Application `json:"application"`
	Date        *time.Time  `json:"date"`
	DayCount    *int        `json:"day_count"`
	IsActive    *bool       `json:"is_active"`
	Description *string     `json:"description"`
}
