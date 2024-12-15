package service

import (
	"context"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
)

type SpecializationService struct {
	db *database.Database
}

func NewSpecializationService(db *database.Database) *SpecializationService {
	return &SpecializationService{db: db}
}

func (s *SpecializationService) GetSpecializationById(id string) (models.Specialization, error) {
	var specialization models.Specialization

	return specialization, nil
}

func (s *SpecializationService) GetAllSpecializations() ([]models.Specialization, error) {
	var specializations []models.Specialization

	query := `SELECT id, name FROM specialization`

	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return specializations, err
	}
	defer rows.Close()

	for rows.Next() {
		var specialization models.Specialization
		err := rows.Scan(&specialization.ID, &specialization.Name)
		if err != nil {
			return specializations, err
		}
		specializations = append(specializations, specialization)
	}

	return specializations, nil
}
