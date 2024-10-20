package transport

import (
	"time"

	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(service *service.Service) *gin.Engine {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userHandler := NewUserHandler(service.UserService)
	appealsHandler := NewAppealHandler(service.AppealService)
	applicationsHandler := NewApplicationHandler(service.ApplicationService)

	auth := router.Group("/auth")
	{
		auth.POST("/login", userHandler.LoginHandler)
		auth.POST("/register", userHandler.RegistrationHandler)
	}

	api := router.Group("/api", AuthMiddleware())
	{
		appeals := api.Group("/appeals")
		{
			appeals.POST("", appealsHandler.CreateAppealHandler)
			appeals.GET("", appealsHandler.GetAllAppealsHandler)
			appeals.GET("/:id", appealsHandler.GetAppealHandler)
			appeals.PUT("", appealsHandler.UpdateAppealHandler)
			appeals.DELETE("/:id", appealsHandler.DeleteAppealHandler)
		}

		applications := api.Group("/applications")
		{
			applications.POST("", applicationsHandler.CreateApplicationHandler)
			applications.GET("", applicationsHandler.GetAllApplicationsHandler)
			applications.GET("/:id", applicationsHandler.GetApplicationHandler)
			applications.PUT("", applicationsHandler.UpdateApplicationHandler)
			applications.DELETE("/:id", applicationsHandler.DeleteApplicationHandler)
		}
	}

	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Status(204)
	})

	return router
}
