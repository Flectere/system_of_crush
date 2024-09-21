package service

import (
	"context"
	"errors"
	"time"

	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const salt = "z9pcWW2o19sTcBw7V075"

type UserService struct {
	db *database.Database
}

func generateJWT(user models.User) (string, error) {
	secretKey := []byte("emagertlESPi")

	claims := jwt.MapClaims{
		"ID":         user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"patronymic": user.Patronymic,
		"role_id":    user.RoleID,
		"exp":        time.Now().Add(time.Hour * 12).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", errors.New("ошибка создания токена")
	}

	return tokenString, nil
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
func (s *UserService) Login(username, password string) (string, error) {
	var user models.User
	err := s.db.Pool.QueryRow(context.Background(), `select * from "user" where login = $1`, username).Scan(&user.ID, &user.Login, &user.Password, &user.LastName, &user.FirstName, &user.Patronymic, &user.RoleID)
	if err != nil {
		return "", errors.New("пользователь не найден")
	}

	err = checkPassword(password, user.Password)
	if err != nil {
		return "", errors.New("пароль неверный")
	}

	token, err := generateJWT(user)
	if err != nil {
		return "", errors.New("ошибка создания токена")
	}

	return token, nil
}

// Создание нового пользователя в базе данных
func (s *UserService) CreateUser(user models.User) (int, error) {
	var id int

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return 0, errors.New("не удалось захэшировать пароль")
	}
	user.Password = hashedPassword

	row := s.db.Pool.QueryRow(context.Background(), `INSERT INTO "user" (login, password, last_name, first_name, patronymic, id_role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID`, user.Login, user.Password, user.LastName, user.FirstName, user.Patronymic, user.RoleID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
