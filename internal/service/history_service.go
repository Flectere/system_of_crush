package service

import (
	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
	"golang.org/x/net/context"
)

type HistoryService struct {
	db *database.Database
}

func NewHistoryService(db *database.Database) *HistoryService {
	return &HistoryService{db: db}
}

func (s *HistoryService) GetAppeals() ([]models.AppealList, error) {
	var appeals []models.AppealList
	query := `
		SELECT ap.id, ap.create_date, con."name", ap.address
		FROM appeal ap
		LEFT JOIN accident_content con ON ap.id_accident = con.id
		LEFT JOIN accident_character ac ON con.id_character = ac.id
		LEFT JOIN specialization spec ON ac.id_specialization = spec.id  
		WHERE ap.is_active = false
		ORDER BY ap.id
		`
	disableAppeal(s.db)

	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var appeal models.AppealList

		err := rows.Scan(
			&appeal.ID,
			&appeal.CreateDate,
			&appeal.Accident,
			&appeal.Address,
		)

		if err != nil {
			return nil, err
		}

		appeals = append(appeals, appeal)
	}

	return appeals, nil
}

func (s *HistoryService) GetShutdowns() ([]models.ShutdownList, error) {
	var allShutdowns []models.ShutdownList
	query :=
		`
		SELECT shut.id, shut.address, shut.date, shut.day_count, con."name"
		FROM shutdown shut
		LEFT JOIN accident_content con ON shut.id_accident = con.id
		LEFT JOIN accident_character char ON con.id_character = char.id
		LEFT JOIN specialization spec ON char.id_specialization = spec.id
		WHERE shut.is_active = false
		ORDER BY shut.id
	`
	err := disableShutdowns(s.db)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var shutdown models.ShutdownList

		err := rows.Scan(
			&shutdown.ID,
			&shutdown.Address,
			&shutdown.Date,
			&shutdown.DayCount,
			&shutdown.Accident,
		)
		if err != nil {
			return nil, err
		}

		allShutdowns = append(allShutdowns, shutdown)
	}

	return allShutdowns, nil
}
func (s *HistoryService) GetApplications() ([]models.ApplicationList, error) {
	var allApplications []models.ApplicationList

	query := `SELECT 
				ap.id, ap.create_date, con."name", im."name", ap.address
			FROM application ap
				LEFT JOIN status st ON ap.id_status = st.id
				LEFT JOIN importance im ON ap.id_importance = im.id
				LEFT JOIN accident_content con ON ap.id_accident = con.id
				JOIN accident_character char ON con.id_character = char.id
				JOIN specialization spec ON char.id_specialization = spec.id
			WHERE id_status = 4
			ORDER BY ap.id
	`

	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var application models.ApplicationList

		err := rows.Scan(
			&application.ID,
			&application.CreateDate,
			&application.Accident,
			&application.Importance,
			&application.Address,
		)
		if err != nil {
			return nil, err
		}

		allApplications = append(allApplications, application)
	}

	return allApplications, nil
}
