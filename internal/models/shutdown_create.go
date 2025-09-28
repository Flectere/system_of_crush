package models

import "time"

type ShutdownCreate struct {
	Address     *string    `json:"address"`
	Accident    *int       `json:"id_accident"`
	Date        *time.Time `json:"date"`
	DayCount    *int       `json:"day_count"`
	HoursCount  *int       `json:"hours_count"`
	Application *int       `json:"id_application"`
	Description *string    `json:"description"`
	Type        *int       `json:"id_type"`
}
