package service

import (
	"context"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
)

type StatusService struct {
	db *database.Database
}

func NewStatusService(db *database.Database) *StatusService {
	return &StatusService{db: db}
}

func (s *StatusService) GetStatusById(id string) (models.Status, error) {
	// Выполнить запрос к БД и вернуть информацию об аварии
	return models.Status{}, nil
}

func (s *StatusService) GetAllStatuses() ([]models.Status, error) {
	var statuses []models.Status
	query := `SELECT id, name FROM status`

	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var status models.Status

		err := rows.Scan(
			&status.ID,
			&status.Name,
		)
		if err != nil {
			return nil, err
		}

		statuses = append(statuses, status)
	}

	return statuses, nil
}
