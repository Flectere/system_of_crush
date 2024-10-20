package service

import (
	"context"

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
		(description, create_date, id_accident, id_importance, id_status, applicant_name, applicant_number)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
	`

	err := s.db.Pool.QueryRow(context.Background(), query,
		appeal.Description,
		appeal.CreateDate,
		appeal.Accident.ID,
		appeal.Importance.ID,
		appeal.Status.ID,
		appeal.ApplicantName,
		appeal.ApplicantNumber,
	).Scan(&newAppealID)

	if err != nil {
		return 0, err
	}

	return newAppealID, nil
}

// Получение всех обращений
func (s *AppealService) GetAllAppeals() ([]models.Appeal, error) {
	var allModels []models.Appeal
	query := `SELECT ap.id, ap.create_date, spec."name", con."name", im."name",st."name"
				FROM appeal ap
				JOIN status st ON ap.id_status = st.id
				JOIN importance im ON ap.id_importance = im.id
				JOIN accident_content con ON ap.id_accident = con.id
				JOIN accident_character char ON con.id_character = char.id
				JOIN specialization spec ON char.id_specialization = spec.id
				ORDER BY ap.id
	`
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
			&appeal.Importance.Name,
			&appeal.Status.Name,
		)
		if err != nil {
			return nil, err
		}

		allModels = append(allModels, appeal)
	}

	return allModels, nil
}

// Получение обращения по id
func (s *AppealService) GetAppealById(id string) (models.Appeal, error) {
	var appeal models.Appeal

	query := `
		SELECT ap.id, ap.create_date, ap.description, st."name", im."name", con."name", char."name", spec."name", ap.applicant_name, ap.applicant_number 
		FROM appeal ap
		JOIN status st ON ap.id_status = st.id
		JOIN importance im ON ap.id_importance = im.id
		JOIN accident_content con ON ap.id_accident = con.id
		JOIN accident_character char ON con.id_character = char.id
		JOIN specialization spec ON char.id_specialization = spec.id
		WHERE ap.id = $1
	`
	row := s.db.Pool.QueryRow(context.Background(), query, id)

	err := row.Scan(
		&appeal.ID,
		&appeal.CreateDate,
		&appeal.Description,
		&appeal.Status.Name,
		&appeal.Importance.Name,
		&appeal.Accident.Name,
		&appeal.Accident.Character.Name,
		&appeal.Accident.Character.Specialization.Name,
		&appeal.ApplicantName,
		&appeal.ApplicantNumber,
	)
	if err != nil {
		return appeal, err
	}

	return appeal, nil
}

func (s *AppealService) UpdateAppeal(appeal models.Appeal) error {
	query := `UPDATE appeal
	SET id=$1, description=$2, create_date=$3, id_accident=$4, id_importance=$5, id_status=$6, applicant_name=$7, applicant_number=$8
	WHERE id=$1;`

	_, err := s.db.Pool.Exec(context.Background(), query, appeal.ID, appeal.Description, appeal.CreateDate, appeal.Accident.ID, appeal.Importance.ID, appeal.Status.ID, appeal.ApplicantName, appeal.ApplicantNumber)

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
