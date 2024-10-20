package service

import (
	"context"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
)

type ApplicationService struct {
	db *database.Database
}

func NewApplicationService(db *database.Database) *ApplicationService {
	return &ApplicationService{db: db}
}

// Создание обращения
func (s *ApplicationService) CreateApplication(application models.Application) (int, error) {
	var newapplicationID int

	query := `
		INSERT INTO application
		(description, create_date, id_accident, id_importance, id_status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
	`

	err := s.db.Pool.QueryRow(context.Background(), query,
		application.Description,
		application.CreateDate,
		application.Accident.ID,
		application.Importance.ID,
		application.Status.ID,
	).Scan(&newapplicationID)

	if err != nil {
		return 0, err
	}

	return newapplicationID, nil
}

// Получение всех обращений
func (s *ApplicationService) GetAllApplications() ([]models.Application, error) {
	var allModels []models.Application

	query := `SELECT ap.id, ap.create_date, spec."name", con."name", im."name",st."name"
				FROM application ap
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
		var application models.Application

		err := rows.Scan(
			&application.ID,
			&application.CreateDate,
			&application.Accident.Character.Specialization.Name,
			&application.Accident.Name,
			&application.Importance.Name,
			&application.Status.Name,
		)
		if err != nil {
			return nil, err
		}

		allModels = append(allModels, application)
	}

	return allModels, nil
}

// Получение обращения по id
func (s *ApplicationService) GetApplicationById(id string) (models.Application, error) {
	var application models.Application

	query := `
		SELECT ap.id, ap.create_date, ap.description, st."name", im."name", con."name", char."name", spec."name"
		FROM application ap
		JOIN status st ON ap.id_status = st.id
		JOIN importance im ON ap.id_importance = im.id
		JOIN accident_content con ON ap.id_accident = con.id
		JOIN accident_character char ON con.id_character = char.id
		JOIN specialization spec ON char.id_specialization = spec.id
		WHERE ap.id = $1
	`
	row := s.db.Pool.QueryRow(context.Background(), query, id)

	err := row.Scan(
		&application.ID,
		&application.CreateDate,
		&application.Description,
		&application.Status.Name,
		&application.Importance.Name,
		&application.Accident.Name,
		&application.Accident.Character.Name,
		&application.Accident.Character.Specialization.Name,
	)
	if err != nil {
		return application, err
	}

	return application, nil
}

func (s *ApplicationService) UpdateApplication(application models.Application) error {
	query := `UPDATE application
	SET id=$1, description=$2, create_date=$3, id_accident=$4, id_importance=$5, id_status=$6
	WHERE id=$1;`

	_, err := s.db.Pool.Exec(context.Background(), query, application.ID, application.Description, application.CreateDate, application.Accident.ID, application.Importance.ID, application.Status.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *ApplicationService) DeleteApplication(id string) error {
	query := `DELETE FROM application WHERE id=$1;`

	_, err := s.db.Pool.Exec(context.Background(), query, id)

	if err != nil {
		return err
	}

	return nil
}
