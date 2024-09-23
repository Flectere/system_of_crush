package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Flectere/system_of_crush/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "emagertlESPifsdjhhfasduBULATAXAXAXAXA"

// Генерация JWT токена
func generateJWT(user models.User) (string, error) {

	claims := jwt.MapClaims{
		"ID":         user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"patronymic": user.Patronymic,
		"role_id":    user.RoleID,
		"exp":        time.Now().Add(time.Hour * 12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.New("ошибка создания токена")
	}

	return tokenString, nil
}

// Валидация JWT токена
func ValidateJWT(tokenString string) (int, error) {
	claims := jwt.MapClaims{}
	// Проверка того ,что токен был создан нашим алгоритмом с использованием нашего ключа
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	// Извлечение ID из claims токена
	id, ok := claims["ID"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid ID in claims")
	}

	return int(id), nil
}
