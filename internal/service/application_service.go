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

// Создание заявки
func (s *ApplicationService) CreateApplication(application models.Application) (int, error) {
	var newapplicationID int

	query := `
		INSERT INTO application
		(description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id
	`

	err := s.db.Pool.QueryRow(context.Background(), query,
		application.Description,
		application.CreateDate,
		application.Accident.ID,
		application.Importance.ID,
		application.Status.ID,
		application.Address,
		application.AccidentCause,
		application.DamagePoint,
		application.Material.ID,
		application.DamageType.ID,
	).Scan(&newapplicationID)

	if err != nil {
		return 0, err
	}

	return newapplicationID, nil
}

// Получение всех заявок
func (s *ApplicationService) GetAllApplications() ([]models.Application, error) {
	var allApplications []models.Application

	query := `SELECT ap.id, ap.create_date, spec."name", con."name", im."name", st."name", ap.address
				FROM application ap
				LEFT JOIN status st ON ap.id_status = st.id
				LEFT JOIN importance im ON ap.id_importance = im.id
				LEFT JOIN accident_content con ON ap.id_accident = con.id
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
			&application.Address,
		)
		if err != nil {
			return nil, err
		}

		allApplications = append(allApplications, application)
	}

	return allApplications, nil
}

// Получение всех заявок у бригадира
func (s *ApplicationService) GetAllBrigadirApplications(id string) ([]models.Application, error) {
	var allApplications []models.Application

	query := `SELECT ap.id, ap.create_date, spec."name", con."name", im."name", st."name", ap.address
				FROM application ap
				LEFT JOIN status st ON ap.id_status = st.id
				LEFT JOIN importance im ON ap.id_importance = im.id
				LEFT JOIN accident_content con ON ap.id_accident = con.id
				JOIN accident_character char ON con.id_character = char.id
				JOIN specialization spec ON char.id_specialization = spec.id
				JOIN brigade ON brigade.id = ap.id_brigade
				JOIN "user" brigadir ON brigadir.id = brigade.id_brigadir
				WHERE brigadir.id = 2
				ORDER BY ap.id
	`

	rows, err := s.db.Pool.Query(context.Background(), query, id)
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

// Получение заявки по id
func (s *ApplicationService) GetApplicationById(id string) (models.Application, error) {
	var application models.Application

	query := `
		SELECT ap.id, ap.description, ap.create_date, st.ID, st."name", im.ID, im."name", con.id, con."name", char.id, char."name", spec.id, spec."name", ap.address, dm.id, dm."name", mat.id, mat."name", ap.accident_cause, ap.damage_point, br.id, br.people_count, us.id, us.last_name, us.first_name, us.patronymic
			FROM application ap
			LEFT JOIN damage_type dm ON ap.id_damage = dm.id
			LEFT JOIN material_type mat ON ap.id_material = mat.id
			LEFT JOIN status st ON ap.id_status = st.id
			LEFT JOIN importance im ON ap.id_importance = im.id
			LEFT JOIN accident_content con ON ap.id_accident = con.id
			LEFT JOIN brigade br ON ap.id_brigade = br.id
			LEFT JOIN "user" us ON br.id_brigadir = us.id
			JOIN accident_character char ON con.id_character = char.id
			JOIN specialization spec ON char.id_specialization = spec.id
			WHERE ap.id = $1
	`
	row := s.db.Pool.QueryRow(context.Background(), query, id)
	err := row.Scan(
		&application.ID,
		&application.Description,
		&application.CreateDate,
		&application.Status.ID,
		&application.Status.Name,
		&application.Importance.ID,
		&application.Importance.Name,
		&application.Accident.ID,
		&application.Accident.Name,
		&application.Accident.Character.ID,
		&application.Accident.Character.Name,
		&application.Accident.Character.Specialization.ID,
		&application.Accident.Character.Specialization.Name,
		&application.Address,
		&application.DamageType.ID,
		&application.DamageType.Name,
		&application.Material.ID,
		&application.Material.Name,
		&application.AccidentCause,
		&application.DamagePoint,
		&application.Brigade.ID,
		&application.Brigade.PeopleCount,
		&application.Brigade.Brigadir.ID,
		&application.Brigade.Brigadir.LastName,
		&application.Brigade.Brigadir.FirstName,
		&application.Brigade.Brigadir.Patronymic,
	)
	if err != nil {
		return application, err
	}

	return application, nil
}

// Изменение заявки
func (s *ApplicationService) UpdateApplication(application models.Application) error {
	query := `UPDATE application
	SET id=$1, description=$2, create_date=$3, id_accident=$4, id_importance=$5, id_status=$6, address=$7, accident_cause=$8, damage_point=$9, id_material=$10, id_damage=$11
	WHERE id=$1;`
	_, err := s.db.Pool.Exec(context.Background(), query, application.ID, application.Description, application.CreateDate, application.Accident.ID, application.Importance.ID, application.Status.ID, application.Address, application.AccidentCause, application.DamagePoint, application.Material.ID, application.DamageType.ID)

	if err != nil {
		return err
	}

	return nil
}

// Удаление заявки
func (s *ApplicationService) DeleteApplication(id string) error {
	query := `DELETE FROM application WHERE id=$1;`

	_, err := s.db.Pool.Exec(context.Background(), query, id)

	if err != nil {
		return err
	}

	return nil
}
