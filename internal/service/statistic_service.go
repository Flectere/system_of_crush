package service

import (
	"context"
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

func (s *StatisticService) GetStatistics() models.Statistic {
	var statistic models.Statistic

	now := time.Now()
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, 0)

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
