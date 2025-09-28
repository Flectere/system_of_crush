package models

type Statistic struct {
	Appeals             *int `json:"appeals"`
	Applications        *int `json:"applications"`
	Shutdowns           *int `json:"shutdowns"`
	AppealsApplications *int `json:"appeals_applications"` // Обращения к которым прикреплена заявка
}
