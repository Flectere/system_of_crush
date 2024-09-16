package service

import (
	"context"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
)

type UserService struct {
	db *database.Database
}

func NewUserService(db *database.Database) *UserService {
	return &UserService{db: db}
}

// Создание нового пользователя в базе данных
func (s *UserService) CreateUser(user models.User) error {
	_, err := s.db.Pool.Exec(context.Background(), `INSERT INTO "user" (login, password, last_name, first_name, patronymic, id_role) VALUES ($1, $2, $3, $4, $5, $6)`, user.Login, user.Password, user.LastName, user.FirstName, user.Password, user.RoleID)
	return err
}
