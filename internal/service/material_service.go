package service

import (
	"context"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
)

type MaterialService struct {
	db *database.Database
}

func NewMaterialService(db *database.Database) *MaterialService {
	return &MaterialService{db: db}
}

func (s *MaterialService) GetAllMaterials() ([]models.Material, error) {
	var materials []models.Material

	query := `SELECT id, name FROM material_type`

	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var material models.Material

		err := rows.Scan(
			&material.ID,
			&material.Name,
		)
		if err != nil {
			return nil, err
		}

		materials = append(materials, material)
	}

	return materials, nil
}
