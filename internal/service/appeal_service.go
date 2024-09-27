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

	// SQL-запрос для вставки нового обращения
	query := `
		INSERT INTO appeal (title, description, create_date, id_specialization, id_importance) 
		VALUES ($1, $2, NOW(), $3, $4)
		RETURNING id
	`

	// Выполнение запроса и получение ID нового обращения
	err := s.db.Pool.QueryRow(context.Background(), query,
		appeal.Title,
		appeal.Description,
		appeal.Specialization.ID,
		appeal.Importance.ID,
	).Scan(&newAppealID)

	if err != nil {
		return 0, err
	}

	return newAppealID, nil
}

// Получение всех обращений
func (s *AppealService) GetAllAppeals() ([]models.Appeal, error) {
	var allModels []models.Appeal

	// SQL запрос для получения данных о всех обращениях
	query := `SELECT a.id, a.title, s.id, s.name, i.id, i.name 
	FROM "appeal" a
	JOIN specialization s ON a.id_specialization = s.id
	JOIN importance i ON a.id_importance = i.id
	`
	// Выполнение запроса
	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Заполнение списка обращениями
	for rows.Next() {
		var appeal models.Appeal

		err := rows.Scan(
			&appeal.ID,
			&appeal.Title,
			&appeal.Specialization.ID,
			&appeal.Specialization.Name,
			&appeal.Importance.ID,
			&appeal.Importance.Name,
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

	// SQL-запрос для получения данных об обращении
	query := `
		SELECT a.id, a.title, a.description, a.create_date, s.id, s.name, i.id, i.name 
		FROM "appeal" a
		JOIN specialization s ON a.id_specialization = s.id
		JOIN importance i ON a.id_importance = i.id
		WHERE a.id = $1
	`
	// Выполняем запрос
	row := s.db.Pool.QueryRow(context.Background(), query, id)

	// Сканируем данные в структуру Appeal, включая вложенные структуры
	err := row.Scan(
		&appeal.ID,
		&appeal.Title,
		&appeal.Description,
		&appeal.CreateDate,
		&appeal.Specialization.ID,
		&appeal.Specialization.Name,
		&appeal.Importance.ID,
		&appeal.Importance.Name,
	)
	if err != nil {
		return appeal, err
	}

	return appeal, nil
}
