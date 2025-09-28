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
func (s *AppealService) CreateAppeal(appeal models.AppealCreate) (int, error) {
	var newAppealID int

	query := `
		INSERT INTO appeal
		(description, create_date, id_accident, applicant_name, applicant_number, additional_number, address, id_application)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
	`

	err := s.db.Pool.QueryRow(context.Background(), query,
		appeal.Description,
		appeal.CreateDate,
		appeal.Accident,
		appeal.ApplicantName,
		appeal.ApplicantNumber,
		appeal.AdditionalNumber,
		appeal.Address,
		appeal.Application,
	).Scan(&newAppealID)

	if err != nil {
		return 0, err
	}

	return newAppealID, nil
}

func disableAppeal(db *database.Database) error {
	query := `
			UPDATE appeal
				SET is_active = false
			FROM application
			WHERE 
				(appeal.id_application IS NULL AND appeal.create_date <= $1)
				OR 
				(appeal.id_application = application.ID AND application.id_status = 4);

`
	sixMonthsAgo := time.Now().AddDate(0, -6, 0)

	_, err := db.Pool.Exec(context.Background(), query, sixMonthsAgo)
	return err
}

// Получение всех обращений
func (s *AppealService) GetAllAppeals() ([]models.AppealList, error) {
	var allApeals []models.AppealList
	query := `SELECT ap.id, ap.create_date, con."name", ap.address, spec.name,
    CASE  
        WHEN ap.id_application IS NULL THEN false  
        ELSE true  
    END AS status  
	FROM appeal ap 
	JOIN accident_content con ON ap.id_accident = con.id  
	JOIN accident_character ac ON con.id_character = ac.id
	JOIN specialization spec ON ac.id_specialization = spec.id  
	WHERE ap.is_active = true
	ORDER BY ap.id;
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
		var appeal models.AppealList

		err := rows.Scan(
			&appeal.ID,
			&appeal.CreateDate,
			&appeal.Accident,
			&appeal.Address,
			&appeal.Specialization,
			&appeal.Status,
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

func (s *AppealService) UpdateAppeal(appeal models.AppealEdit) error {
	query := `
				UPDATE appeal
				SET id=$1, description=$2, create_date=$3, id_accident=$4, applicant_name=$5, applicant_number=$6, additional_number=$7, address=$8, id_application=$9
				WHERE id=$1;
				`

	_, err := s.db.Pool.Exec(context.Background(), query, appeal.ID, appeal.Description, appeal.CreateDate, appeal.Accident, appeal.ApplicantName, appeal.ApplicantNumber, appeal.AdditionalNumber, appeal.Address, appeal.Application)

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
