package models

type Brigade struct {
	ID          *int `json:"id,omitempty"`
	PeopleCount *int `json:"people_count,omitempty"`
	Brigadir    User `json:"brigadir,omitempty"`
}