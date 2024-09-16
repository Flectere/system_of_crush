package transport

import (
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

func NewRouter(userService *service.UserService) *gin.Engine {
	router := gin.Default()

	userhandler := NewUserHandler(userService)

	auth := router.Group("/auth")
	{
		auth.POST("/login")
		auth.POST("/register", userhandler.CreateUser)
	}

	return router
}
