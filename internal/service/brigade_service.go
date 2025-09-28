package service

import (
	"context"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
)

type BrigadeService struct {
	db *database.Database
}

func NewBrigadeService(db *database.Database) *BrigadeService {
	return &BrigadeService{db}
}

func (s *BrigadeService) GetFreeBrigadirs() ([]models.Brigadir, error) {
	var brigadirs []models.Brigadir
	query := `
		SELECT 
			u.id, u.last_name, u.first_name, u.patronymic
		FROM "user" u
			LEFT JOIN brigade b ON b.id_brigadir = u.id
		WHERE b.id_brigadir IS NULL AND u.id_role = 3	
		`
	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var brigadir models.Brigadir
		err := rows.Scan(
			&brigadir.ID,
			&brigadir.LastName,
			&brigadir.FirstName,
			&brigadir.Patronymic,
		)
		if err != nil {
			return nil, err
		}

		brigadirs = append(brigadirs, brigadir)
	}
	return brigadirs, nil
}

func (s *BrigadeService) GetAllBrigades() ([]models.BrigadeList, error) {
	var brigades []models.BrigadeList
	query := `
		SELECT 
			b.id, b.people_count, u.last_name, u.first_name, u.patronymic
		FROM brigade b
			JOIN "user" u ON b.id_brigadir = u.id	
	`

	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return brigades, err
	}
	defer rows.Close()

	for rows.Next() {
		var brigade models.BrigadeList
		err := rows.Scan(
			&brigade.ID,
			&brigade.PeopleCount,
			&brigade.LastName,
			&brigade.FirstName,
			&brigade.Patronymic,
		)
		if err != nil {
			return brigades, err
		}

		brigades = append(brigades, brigade)
	}

	return brigades, nil
}

func (s *BrigadeService) CreateBrigade(brigade models.BrigadeCreate) (int, error) {
	var newBrigadeId int

	query := `
		INSERT INTO brigade 
			(people_count, id_brigadir)
			VALUES ($1,$2)
		RETURNING id
	`
	row := s.db.Pool.QueryRow(context.Background(), query,
		brigade.PeopleCount,
		brigade.Brigadir,
	)

	err := row.Scan(&newBrigadeId)
	if err != nil {
		return 0, err
	}

	return newBrigadeId, nil
}

func (s *BrigadeService) EditBrigade(id string, brigade models.BrigadeCreate)  error {
	query := `
		UPDATE brigade 
		SET 
			people_count = $1,
			id_brigadir = $2
		WHERE id = $3
	`
	_, err := s.db.Pool.Exec(context.Background(), query, brigade.PeopleCount, brigade.Brigadir, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *BrigadeService) GetBrigadeByID(id string) (models.Brigade, error) {
	var brigade models.Brigade
	query := `
		SELECT 
			b.id, b.people_count, u.last_name, u.first_name, u.patronymic, u.id
		FROM brigade b
			JOIN "user" u ON b.id_brigadir = u.id
		WHERE b.id = $1
	`
	row := s.db.Pool.QueryRow(context.Background(), query, id)
	row.Scan(
		&brigade.ID,
		&brigade.PeopleCount,
		&brigade.LastName,
		&brigade.FirstName,
		&brigade.Patronymic,
		&brigade.BrigadirID,
	)

	return brigade, nil
}
