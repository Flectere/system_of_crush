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
func (s *ApplicationService) CreateApplication(application models.ApplicationCreate) (int, error) {
	var newapplicationID int

	query := `
		INSERT INTO application
		(description, create_date, id_accident, id_importance, id_status, address, accident_cause, id_material, id_damage, damage_point, id_operator)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id
	`

	err := s.db.Pool.QueryRow(context.Background(), query,
		application.Description,
		application.CreateDate,
		application.Accident,
		application.Importance,
		application.Status,
		application.Address,
		application.AccidentCause,
		application.Material,
		application.DamageType,
		application.DamagePoint,
		application.Operator,
	).Scan(&newapplicationID)

	if err != nil {
		return 0, err
	}

	return newapplicationID, nil
}

func (s *ApplicationService) GetAllFreeApplications() ([]models.ApplicationList, error) {
	var allApplications []models.ApplicationList

	query := `SELECT 
				ap.id, ap.create_date, con."name", im."name", st."name", ap.address, spec.name
			FROM application ap
				LEFT JOIN status st ON ap.id_status = st.id
				LEFT JOIN importance im ON ap.id_importance = im.id
				LEFT JOIN accident_content con ON ap.id_accident = con.id
				JOIN accident_character char ON con.id_character = char.id
				JOIN specialization spec ON char.id_specialization = spec.id
			WHERE ap.id_status = 1
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
			&application.Status,
			&application.Address,
			&application.Specialization,
		)
		if err != nil {
			return nil, err
		}

		allApplications = append(allApplications, application)
	}

	return allApplications, nil
}

// Получение всех заявок
func (s *ApplicationService) GetAllApplications() ([]models.ApplicationList, error) {
	var allApplications []models.ApplicationList

	query := `SELECT ap.id, ap.create_date, con."name", im."name", st."name", ap.address, spec.name
				FROM application ap
				LEFT JOIN status st ON ap.id_status = st.id
				LEFT JOIN importance im ON ap.id_importance = im.id
				LEFT JOIN accident_content con ON ap.id_accident = con.id
				JOIN accident_character char ON con.id_character = char.id
				JOIN specialization spec ON char.id_specialization = spec.id
				WHERE ap.id_status != 3
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
			&application.Status,
			&application.Address,
			&application.Specialization,
		)
		if err != nil {
			return nil, err
		}

		allApplications = append(allApplications, application)
	}

	return allApplications, nil
}

// Получение всех заявок у бригадира
func (s *ApplicationService) GetAllBrigadirApplications(id string) ([]models.ApplicationList, error) {
	var allApplications []models.ApplicationList
	query := `
			SELECT ap.id, con."name", im."name", ap.address, st.name
			FROM application ap
				LEFT JOIN importance im ON ap.id_importance = im.id
				LEFT JOIN accident_content con ON ap.id_accident = con.id
				JOIN brigade ON brigade.id = ap.id_brigade
				JOIN "user" brigadir ON brigadir.id = brigade.id_brigadir
				JOIN status st ON st.id = ap.id_status 
			WHERE brigadir.id = $1 AND (ap.id_status = 2 OR ap.id_status = 3)
			ORDER BY ap.id
	`

	rows, err := s.db.Pool.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var application models.ApplicationList

		err := rows.Scan(
			&application.ID,
			&application.Accident,
			&application.Importance,
			&application.Address,
			&application.Status,
		)
		if err != nil {
			return nil, err
		}

		allApplications = append(allApplications, application)
	}

	return allApplications, nil
}

func (s *ApplicationService) GetApllicationByIdMobile(id string) (models.ApplicationBrigadir, error) {
	var application models.ApplicationBrigadir

	query := `
		SELECT 
			ap.id, ap.address, con."name", im."name", char.name, mat.name, dam.name, ap.damage_point, ap.description, ap.start_date, con.recommended_time
		FROM application ap
			LEFT JOIN accident_content con ON ap.id_accident = con.id
			LEFT JOIN importance im ON ap.id_importance = im.id
			LEFT JOIN accident_character char ON con.id_character = char.id
			LEFT JOIN material_type mat ON ap.id_material = mat.id
			LEFT JOIN damage_type dam ON ap.id_damage = dam.id
		WHERE ap.id = $1
	`
	row := s.db.Pool.QueryRow(context.Background(), query, id)
	err := row.Scan(
		&application.ID,
		&application.Address,
		&application.Accident,
		&application.Importance,
		&application.Character,
		&application.Material,
		&application.DamageType,
		&application.DamagePoint,
		&application.Description,
		&application.StartDate,
		&application.RecommendedTime,
	)
	if err != nil {
		return application, err
	}

	return application, nil
}

func (s *ApplicationService) SetApplicationToBrigadir(brigadirID, applicationID string) error {
	query := `
		UPDATE application
		SET 
			id_brigade = (
			SELECT id
			FROM brigade
			WHERE id_brigadir = $1
			),
			id_status = 2
		WHERE id = $2
	`
	_, err := s.db.Pool.Exec(context.Background(), query, brigadirID, applicationID)

	return err
}

// Получение заявки по id
func (s *ApplicationService) GetApplicationById(id string) (models.Application, error) {
	var application models.Application

	query := `
		SELECT 
			ap.id, ap.description, ap.create_date, st.ID, st."name", im.ID, im."name", con.id, con."name", char.id, char."name", spec.id, spec."name", ap.address, dm.id, dm."name", mat.id, mat."name", ap.accident_cause, ap.damage_point
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
		// &application.Brigade.ID,
		// &application.Brigade.PeopleCount,
		// &application.Brigade.Brigadir.ID,
		// &application.Brigade.Brigadir.LastName,
		// &application.Brigade.Brigadir.FirstName,
		// &application.Brigade.Brigadir.Patronymic,
	)
	//TODO решить что отправлять в заявке
	if err != nil {
		return application, err
	}

	return application, nil
}

// Изменение заявки
func (s *ApplicationService) UpdateApplication(application models.ApplicationEdit) error {
	query := `
		UPDATE application
		SET id=$1, description=$2, create_date=$3, id_accident=$4, id_importance=$5, id_status=$6, address=$7, accident_cause=$8, id_material=$9, id_damage=$10, damage_point=$11, id_operator = $12  
		WHERE id=$1;`
	_, err := s.db.Pool.Exec(context.Background(), query,
		application.ID,
		application.Description,
		application.CreateDate,
		application.Accident,
		application.Importance,
		application.Status,
		application.Address,
		application.AccidentCause,
		application.Material,
		application.DamageType,
		application.DamagePoint,
		application.Operator,
	)

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

// Старт выполнения заявки
func (s *ApplicationService) StartApplication(id string) error {
	query := `
		UPDATE application
		SET 
			id_status = 3,
			start_date = NOW()
		WHERE id = $1
	`
	_, err := s.db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}

// Завершение выполнения заявки
func (s *ApplicationService) FinishApplication(id string) error {
	query := `
		UPDATE application
		SET 
			id_status = 4,
			finish_date = NOW()
		WHERE id = $1
	`
	_, err := s.db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
