package transport

import (
	"fmt"

	"github.com/Flectere/system_of_crush/internal/models"
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type Userhandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *Userhandler {
	return &Userhandler{userService: userService}
}

func (h *Userhandler) CreateUser(c *gin.Context) {
	//Получение данных о пользователе из запроса
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	fmt.Println(user)

	if err := h.userService.CreateUser(user); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка при создании пользователя %s", err.Error())})
		return
	}

	// Возвращение созданного пользователя в ответе
	c.JSON(201, gin.H{
		"message": "Пользователь успешно создан",
		"user":    user,
	})

}
