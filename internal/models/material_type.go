package models

type Material struct {
	ID   *int    `json:"id"`
	Name *string `json:"name,omitempty"`
}
