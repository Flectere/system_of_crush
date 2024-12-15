package service

import (
	"context"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
)

type DamageService struct {
	db *database.Database
}

func NewDamageService(db *database.Database) *DamageService {
	return &DamageService{db: db}
}

func (s *DamageService) GetAllDamages() ([]models.Damage, error) {
	var damages []models.Damage

	query := `SELECT id, name FROM damage_type`

	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var damage models.Damage

		err := rows.Scan(
			&damage.ID,
			&damage.Name,
		)
		if err != nil {
			return nil, err
		}

		damages = append(damages, damage)
	}

	return damages, nil
}
