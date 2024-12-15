package service

import (
	"context"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
)

type ImportanceService struct {
	db *database.Database
}

func NewImportanceService(db *database.Database) *ImportanceService {
	return &ImportanceService{db: db}
}

func (s *ImportanceService) GetImportanceById(id string) (models.Importance, error) {
	// Выполнить запрос к БД и вернуть информацию об аварии
	return models.Importance{}, nil
}

func (s *ImportanceService) GetAllImportances() ([]models.Importance, error) {
	var importances []models.Importance

	query := `SELECT id, name FROM importance`

	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var importance models.Importance

		err := rows.Scan(
			&importance.ID,
			&importance.Name,
		)
		if err != nil {
			return nil, err
		}

		importances = append(importances, importance)
	}

	return importances, nil
}
