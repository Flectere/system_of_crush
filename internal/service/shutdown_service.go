package service

import (
	"context"
	"time"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
)

type ShutdownService struct {
	db *database.Database
}

func NewShutdownService(db *database.Database) *ShutdownService {
	return &ShutdownService{db: db}
}

func (s *ShutdownService) CreateShutdown(shutdown models.ShutdownCreate) (int, error) {
	var id int
	query := `INSERT INTO shutdown 
				(address, id_accident, date, day_count, hour_count, id_application, description, id_type) 
				VALUES($1, $2, $3, $4, $5, $6, $7, $8) returning id
				`
	row := s.db.Pool.QueryRow(context.Background(), query, shutdown.Address, shutdown.Accident, shutdown.Date, shutdown.DayCount, shutdown.HoursCount, shutdown.Application, shutdown.Description, shutdown.Type)

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func disableShutdowns(db *database.Database) error {
	now := time.Now()
	query := `
		UPDATE shutdown 
		SET is_active = false 
		WHERE date + 
			COALESCE(day_count, 0) * INTERVAL '1 day' + 
			COALESCE(hour_count, 0) * INTERVAL '1 hour' < $1
	`

	_, err := db.Pool.Exec(context.Background(), query, now)
	return err
}

func (s *ShutdownService) GetAllShutdowns() ([]models.ShutdownList, error) {
	var allShutdowns []models.ShutdownList
	query :=
		`
		SELECT shut.id, shut.address, shut.date, shut.day_count, con."name", spec."name", type."name"
		FROM shutdown shut
		LEFT JOIN accident_content con ON shut.id_accident = con.id
		LEFT JOIN accident_character char ON con.id_character = char.id
		LEFT JOIN specialization spec ON char.id_specialization = spec.id
		JOIN shutdown_type type ON shut.id_type = type.id
		WHERE shut.is_active = true
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
			&shutdown.Specialization,
			&shutdown.Type,
		)
		if err != nil {
			return nil, err
		}

		allShutdowns = append(allShutdowns, shutdown)
	}

	return allShutdowns, nil
}

func (s *ShutdownService) GetShutdownById(id string) (models.Shutdown, error) {
	var shutdown models.Shutdown

	query := `
		SELECT shut.id, shut.address, shut.date, con.id, con."name", char.id, char."name", spec.id, spec."name", shut.id_application, shut.day_count, shut.hour_count, shut.is_active, shut.id_type
		FROM shutdown shut
		JOIN accident_content con ON shut.id_accident = con.id
		JOIN accident_character char ON con.id_character = char.id
		JOIN specialization spec ON char.id_specialization = spec.id
		WHERE shut.id = $1
	`

	row := s.db.Pool.QueryRow(context.Background(), query, id)

	err := row.Scan(
		&shutdown.ID,
		&shutdown.Address,
		&shutdown.Date,
		&shutdown.Accident.ID,
		&shutdown.Accident.Name,
		&shutdown.Accident.Character.ID,
		&shutdown.Accident.Character.Name,
		&shutdown.Accident.Character.Specialization.ID,
		&shutdown.Accident.Character.Specialization.Name,
		&shutdown.Application.ID,
		&shutdown.DayCount,
		&shutdown.HoursCount,
		&shutdown.IsActive,
		&shutdown.Type,
	)
	if err != nil {
		return shutdown, err
	}

	return shutdown, nil
}

func (s *ShutdownService) UpdateShutdown(shutdown models.Shutdown) error {
	query := `
				UPDATE shutdown
				SET id=$1, address=$2, id_accident=$3, date=$4, day_count=$5, hour_count=$6, id_type=$7
				WHERE id=$1;
				`

	_, err := s.db.Pool.Exec(context.Background(), query, shutdown.ID, shutdown.Address, shutdown.Accident.ID, shutdown.Date, shutdown.DayCount, shutdown.HoursCount, shutdown.Type)

	if err != nil {
		return err
	}

	return nil
}

func (s *ShutdownService) DeleteShutdown(id string) error {
	return nil
}
