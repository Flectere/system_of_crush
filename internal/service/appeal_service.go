package service

import (
	"context"
	"time"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
)

type AppealService struct {
	db *database.Database
}

func NewAppealService(db *database.Database) *AppealService {
	return &AppealService{db: db}
}

// Создание обращения
func (s *AppealService) CreateAppeal(appeal models.Appeal) (int, error) {
	var newAppealID int

	query := `
		INSERT INTO appeal
		(description, create_date, id_accident, applicant_name, applicant_number, additional_number, address, id_application, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id
	`

	err := s.db.Pool.QueryRow(context.Background(), query,
		appeal.Description,
		appeal.CreateDate,
		appeal.Accident.ID,
		appeal.ApplicantName,
		appeal.ApplicantNumber,
		appeal.AdditionalNumber,
		appeal.Address,
		appeal.Application.ID,
		appeal.IsActive,
	).Scan(&newAppealID)

	if err != nil {
		return 0, err
	}

	return newAppealID, nil
}

func disableAppeal(db *database.Database) error {
	query := `UPDATE appeal
				SET is_active = false
				FROM application
				WHERE 
				(appeal.id_application IS NULL AND appeal.create_date <= $1)
				OR 
				(appeal.id_application = application.ID AND application.id_status = 3);

`
	sixMonthsAgo := time.Now().AddDate(0, -6, 0)

	_, err := db.Pool.Exec(context.Background(), query, sixMonthsAgo)
	return err
}

// Получение всех обращений
func (s *AppealService) GetAllAppeals() ([]models.Appeal, error) {
	var allApeals []models.Appeal
	query := `
				SELECT ap.id, ap.create_date, spec."name", con."name", ap.address, ap.id_application, ap.applicant_number, ap.is_active
				FROM appeal ap
				JOIN accident_content con ON ap.id_accident = con.id
				JOIN accident_character char ON con.id_character = char.id
				JOIN specialization spec ON char.id_specialization = spec.id
				ORDER BY ap.id
	`
	err := disableAppeal(s.db)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Pool.Query(context.Background(), query)
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
			&appeal.IsActive,
		)

		if err != nil {
			return nil, err
		}

		allApeals = append(allApeals, appeal)
	}

	return allApeals, nil
}

// Получение обращения по id
func (s *AppealService) GetAppealById(id string) (models.Appeal, error) {
	var appeal models.Appeal

	query := `
				SELECT ap.id, ap.create_date, ap.description, ap."address", con.id, con."name",char.id, char."name", spec.id, spec."name", ap.applicant_name, ap.applicant_number, ap.additional_number, apl.id, apl."address", apl.id_accident
				FROM appeal ap
				JOIN accident_content con ON ap.id_accident = con.id
				JOIN accident_character char ON con.id_character = char.id
				JOIN specialization spec ON char.id_specialization = spec.id
				LEFT JOIN application apl ON ap.id_application = apl.id
				LEFT JOIN accident_content con_apl ON apl.id_accident = con_apl.id
				WHERE ap.id = $1
	`
	row := s.db.Pool.QueryRow(context.Background(), query, id)

	err := row.Scan(
		&appeal.ID,
		&appeal.CreateDate,
		&appeal.Description,
		&appeal.Address,
		&appeal.Accident.ID,
		&appeal.Accident.Name,
		&appeal.Accident.Character.ID,
		&appeal.Accident.Character.Name,
		&appeal.Accident.Character.Specialization.ID,
		&appeal.Accident.Character.Specialization.Name,
		&appeal.ApplicantName,
		&appeal.ApplicantNumber,
		&appeal.AdditionalNumber,
		&appeal.Application.ID,
		&appeal.Application.Address,
		&appeal.Application.Accident.Name,
	)
	if err != nil {
		return appeal, err
	}

	return appeal, nil
}

func (s *AppealService) UpdateAppeal(appeal models.Appeal) error {
	query := `
				UPDATE appeal
				SET id=$1, description=$2, create_date=$3, id_accident=$4, applicant_name=$5, applicant_number=$6, additional_number=$7, address=$8, id_application=$9
				WHERE id=$1;
				`

	_, err := s.db.Pool.Exec(context.Background(), query, appeal.ID, appeal.Description, appeal.CreateDate, appeal.Accident.ID, appeal.ApplicantName, appeal.ApplicantNumber, appeal.AdditionalNumber, appeal.Address, appeal.Application.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *AppealService) DeleteAppeal(id string) error {
	query := `DELETE FROM appeal WHERE id=$1;`

	_, err := s.db.Pool.Exec(context.Background(), query, id)

	if err != nil {
		return err
	}

	return nil
}
