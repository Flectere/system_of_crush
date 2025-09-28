package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
)

type StatisticService struct {
	db *database.Database
}

func NewStatisticService(db *database.Database) *StatisticService {
	return &StatisticService{db: db}
}

func getCurrentTimeMonth() (time.Time, time.Time) {
	now := time.Now()
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, 0)

	return firstDay, lastDay
}

func (s *StatisticService) GetStatistics() models.Statistic {
	var statistic models.Statistic

	firstDay, lastDay := getCurrentTimeMonth()

	allAppeals := `SELECT COUNT(*) FROM appeal WHERE create_date BETWEEN $1 AND $2;`
	s.db.Pool.QueryRow(context.Background(), allAppeals, firstDay, lastDay).Scan(&statistic.Appeals)

	allApplications := `SELECT COUNT(*) FROM application WHERE create_date BETWEEN $1 AND $2;`
	s.db.Pool.QueryRow(context.Background(), allApplications, firstDay, lastDay).Scan(&statistic.Applications)

	allShutdowns := `SELECT COUNT(*) FROM shutdown WHERE date BETWEEN $1 AND $2;`
	s.db.Pool.QueryRow(context.Background(), allShutdowns, firstDay, lastDay).Scan(&statistic.Shutdowns)

	appealsWithApplication := `SELECT COUNT(*) FROM appeal WHERE id_application IS NOT NULL AND create_date BETWEEN $1 AND $2;`
	s.db.Pool.QueryRow(context.Background(), appealsWithApplication, firstDay, lastDay).Scan(&statistic.AppealsApplications)

	return statistic
}

func (s *StatisticService) GetBrigadirStatistic() ([]models.BrigadirStatistic, error) {
	var statistic []models.BrigadirStatistic
	query := `
			SELECT 
				u.last_name, 
				u.first_name, 
				u.patronymic, 
				COUNT(ap.id), 
				ROUND(EXTRACT(EPOCH FROM AVG(ap.finish_date-ap.start_date)/60)) 
			FROM application ap
				JOIN brigade br ON ap.id_brigade = br.id
				JOIN "user" u ON br.id_brigadir = u.id
			WHERE ap.id_status = 4
			GROUP BY (u.last_name, u.first_name, u.patronymic)
	`

	rows, err := s.db.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var row models.BrigadirStatistic
		var totalMinutes int
		rows.Scan(
			&row.LastName,
			&row.FirstName,
			&row.Patronynic,
			&row.CountApplications,
			&totalMinutes,
		)
		hours := totalMinutes / 60
		minutes := totalMinutes - (hours * 60)

		row.AverageTime = fmt.Sprintf("%d.%d", hours, minutes)

		statistic = append(statistic, row)
	}
	return statistic, nil
}
