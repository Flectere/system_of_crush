package models

type BrigadirStatistic struct {
	LastName          *string `json:"last_name"`
	FirstName         *string `json:"first_name"`
	Patronynic        *string `json:"patronymic"`
	CountApplications *int    `json:"count_applications"`
	AverageTime       string  `json:"average_time"`
}
