package service

import (
	"context"
	"errors"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
	"golang.org/x/crypto/bcrypt"
)

const salt = "z9pcWW2o19sTcBw7V075"

type UserService struct {
	db *database.Database
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), 12)
	return string(hashedPassword), err
}

func checkPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt))

}
func NewUserService(db *database.Database) *UserService {
	return &UserService{db: db}
}

// Аутентификация пользователя по логину и паролю
func (s *UserService) Login(username, password string)

// Создание нового пользователя в базе данных
func (s *UserService) CreateUser(user models.User) (int, error) {
	var id int

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return 0, errors.New("не удалось захэшировать пароль")
	}
	user.Password = hashedPassword

	row := s.db.Pool.QueryRow(context.Background(), `INSERT INTO "user" (login, password, last_name, first_name, patronymic, id_role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID`, user.Login, user.Password, user.LastName, user.FirstName, user.Patromonyc, user.RoleID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
