package service

import (
	"time"

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

func (s *HistoryService) GetAppeals() ([]models.Appeal, error) {
	var appeals []models.Appeal
	query := `
		SELECT ap.id, ap.create_date, spec."name", con."name", ap.address, ap.id_application, ap.applicant_number
		FROM appeal ap
		JOIN accident_content con ON ap.id_accident = con.id
		JOIN accident_character char ON con.id_character = char.id
		JOIN specialization spec ON char.id_specialization = spec.id
		WHERE create_date <= $1
		ORDER BY ap.id
`

	sixMonthsAgo := time.Now().AddDate(0, -6, 0)

	rows, err := s.db.Pool.Query(context.Background(), query, sixMonthsAgo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var appeal models.Appeal

		err := rows.Scan(
			&appeal.ID,
			&appeal.CreateDate,
			&appeal.Accident.Character.Specialization.Name,
			&appeal.Accident.Name,
			&appeal.Address,
			&appeal.Application.ID,
			&appeal.ApplicantNumber,
		)

		if err != nil {
			return nil, err
		}

		appeals = append(appeals, appeal)
	}

	return appeals, nil
}

func (s *HistoryService) GetShutdowns() ([]models.Shutdown, error) {
	var allShutdowns []models.Shutdown
	query :=
		`
		SELECT shut.id, shut.address, shut.date, shut.day_count, shut.is_active, con.id, con."name", char.id, char."name", spec.id, spec."name"
		FROM shutdown shut
		LEFT JOIN accident_content con ON shut.id_accident = con.id
		JOIN accident_character char ON con.id_character = char.id
		JOIN specialization spec ON char.id_specialization = spec.id
		WHERE is_active = false
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
		var shutdown models.Shutdown

		err := rows.Scan(
			&shutdown.ID,
			&shutdown.Address,
			&shutdown.Date,
			&shutdown.DayCount,
			&shutdown.IsActive,
			&shutdown.Accident.ID,
			&shutdown.Accident.Name,
			&shutdown.Accident.Character.ID,
			&shutdown.Accident.Character.Name,
			&shutdown.Accident.Character.Specialization.ID,
			&shutdown.Accident.Character.Specialization.Name,
		)
		if err != nil {
			return nil, err
		}

		allShutdowns = append(allShutdowns, shutdown)
	}

	return allShutdowns, nil
}
func (s *HistoryService) GetApplications() ([]models.Application, error) {
	var allApplications []models.Application

	query := `SELECT ap.id, ap.create_date, spec."name", con."name", im."name", st."name", ap.address
				FROM application ap
				LEFT JOIN status st ON ap.id_status = st.id
				LEFT JOIN importance im ON ap.id_importance = im.id
				LEFT JOIN accident_content con ON ap.id_accident = con.id
				JOIN accident_character char ON con.id_character = char.id
				JOIN specialization spec ON char.id_specialization = spec.id
				WHERE id_status = 3
				ORDER BY ap.id
	`

	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var application models.Application

		err := rows.Scan(
			&application.ID,
			&application.CreateDate,
			&application.Accident.Character.Specialization.Name,
			&application.Accident.Name,
			&application.Importance.Name,
			&application.Status.Name,
			&application.Address,
		)
		if err != nil {
			return nil, err
		}

		allApplications = append(allApplications, application)
	}

	return allApplications, nil
}
