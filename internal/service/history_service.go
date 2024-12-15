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
