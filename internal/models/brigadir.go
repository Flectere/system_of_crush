package models

type Brigadir struct {
	ID         *int    `json:"id"`
	LastName   *string `json:"last_name"`
	FirstName  *string `json:"first_name"`
	Patronymic *string `json:"patronymic"`
}
