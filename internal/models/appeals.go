package models

type Appeals struct {
	ID               int `json:"id"`
	SpecializationID int `json:"id_specialization"`
	ImportanceID     int `json:"id_importance"`
}
