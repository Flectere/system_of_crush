package transport

import (
	"fmt"

	"github.com/Flectere/system_of_crush/internal/models"
	"github.com/Flectere/system_of_crush/internal/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *userHandler {
	return &userHandler{userService: userService}
}

// Обработчик запроса на регистрацию
func (h *userHandler) RegistrationHandler(c *gin.Context) {
	// Получение данных о пользователе из запроса
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		return
	}

	// Добавление пользователя
	id, err := h.userService.Registration(user)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка при создании пользователя %s", err.Error())})
		return
	}

	// Возвращение созданного пользователя в ответе
	c.JSON(201, gin.H{
		"message": "Пользователь успешно создан",
		"user_id": id,
	})
}

// Обработчик запроса на авторизацию
func (h *userHandler) LoginHandler(c *gin.Context) {
	var loginData struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	fmt.Println(c.Request.URL)
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Некорректные данные для входа"})
		return
	}

	token, err := h.userService.Login(loginData.Login, loginData.Password)
	if err != nil {
		switch err.Error() {
		case "пользователь не найден":
			c.JSON(404, gin.H{"error": "Пользователь не найден."})
		case "пароль неверный":
			c.JSON(401, gin.H{"error": "Пароль неверный."})
		default:
			c.JSON(500, gin.H{"error": "Ошибка сервера."})
		}
		return
	}

	c.JSON(200, gin.H{"token": token})
}
