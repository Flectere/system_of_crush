package service

import (
	"context"
	"errors"

	"github.com/Flectere/system_of_crush/config"
	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *database.Database
}

func NewUserService(db *database.Database) *UserService {
	return &UserService{db: db}
}

// Хэширование пароля
func hashPassword(password string) (string, error) {
	salt := config.Config.ServerConfig.Salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), 12)
	return string(hashedPassword), err
}

// Проверка соответсвия пароля и хэшированного пароля из базы
func checkPassword(password, hashedPassword string) error {
	salt := config.Config.ServerConfig.Salt
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt))
}

// Аутентификация пользователя по логину и паролю
func (s *UserService) Login(login, password string) (string, error) {

	user, err := s.getUser(login)
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

// Получение пользователя из базы по логину
func (s *UserService) getUser(login string) (models.User, error) {
	var user models.User

	err := s.db.Pool.QueryRow(context.Background(), `select u.id, u.login, u.password, u.last_name, u.first_name, u.patronymic, r.id, r.name from "user" u JOIN role r ON u.id_role = r.id where u.login = $1;`, login).Scan(&user.ID, &user.Login, &user.Password, &user.LastName, &user.FirstName, &user.Patronymic, &user.Role.ID, &user.Role.Name)
	return user, err
}

// Регистрация пользователя
func (s *UserService) Registration(user models.User) (int, error) {
	var id int

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return 0, errors.New("не удалось захэшировать пароль")
	}
	user.Password = hashedPassword

	row := s.db.Pool.QueryRow(context.Background(), `INSERT INTO "user" (login, password, last_name, first_name, patronymic, id_role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, user.Login, user.Password, user.LastName, user.FirstName, user.Patronymic, user.Role.ID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
