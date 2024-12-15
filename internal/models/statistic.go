package models

type Statistic struct {
	Appeals      int `json:"appeals"`
	Applications int `json:"applications"`
	Shutdowns    int `json:"shutdowns"`
	// Обращения к которым прикреплена заявка
	AppealsApplications int `json:"appeals_applications"`
}
