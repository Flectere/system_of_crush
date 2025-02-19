package service

import (
	"context"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
)

type AccidentService struct {
	db *database.Database
}

func NewAccidentService(db *database.Database) *AccidentService {
	return &AccidentService{db: db}
}

func (s *AccidentService) GetAllAccidents() ([]models.Accident, error) {
	var accidents []models.Accident
	query := `SELECT ac.id, ac.name, char.id, char.name, spec.id, spec.name
				FROM accident_content ac
				JOIN accident_character char ON ac.id_character = char.id
				JOIN specialization spec ON char.id_specialization = spec.id`

	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var accident models.Accident

		err := rows.Scan(
			&accident.ID,
			&accident.Name,
			&accident.Character.ID,
			&accident.Character.Name,
			&accident.Character.Specialization.ID,
			&accident.Character.Specialization.Name,
		)
		if err != nil {
			return nil, err
		}

		accidents = append(accidents, accident)
	}

	return accidents, nil
}
