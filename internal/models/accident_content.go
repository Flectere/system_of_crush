package models

type AccidentContent struct {
	ID        int               `json:"id,omitempty"`
	Name      string            `json:"name,omitempty"`
	Character AccidentCharacter `json:"accident_character"`
}