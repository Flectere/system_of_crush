package models

type BrigadeList struct {
	ID          *int    `json:"id"`
	PeopleCount *int    `json:"people_count"`
	LastName    *string `json:"last_name"`
	FirstName   *string `json:"first_name"`
	Patronymic  *string `json:"patronymic"`
}
