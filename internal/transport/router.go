package transport

import (
	"time"

	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(userService *service.UserService) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userHandler := NewUserHandler(userService)

	// "Эндопоинты для регистрации и авторизации"
	auth := router.Group("/auth")
	{
		auth.POST("/login", userHandler.LoginHandler)
		auth.POST("/register", userHandler.RegistrationHandler)
	}

	router.OPTIONS("/*any", func(c *gin.Context) {
		c.JSON(200, nil)
	})

	return router
}
