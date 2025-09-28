package models

type BrigadeCreate struct {
	PeopleCount *int `json:"people_count,omitempty"`
	Brigadir    *int `json:"id_brigadir,omitempty"`
}
