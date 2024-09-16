package models

type User struct {
	ID         int    `json:"id"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	Patromonyc string `json:"patronymic"`
	RoleID     int    `json:"id_role"`
}
