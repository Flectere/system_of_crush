package transport

import (
	"net/http"
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

	router.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "http://localhost:8084")
	})

	userHandler := newUserHandler(service.UserService)
	appealHandler := newAppealHandler(service.AppealService)
	applicationHandler := newApplicationHandler(service.ApplicationService)
	shutdownHandler := newShutdownHandler(service.ShutdownService)
	accidentHandler := newAccidentHandler(service.AccidentService)
	characterHandler := newCharacterHandler(service.CharacterService)
	specializationHandler := newSpecializationHandler(service.SpecializationService)
	importanceHandler := newImportanceHandler(service.ImportanceService)
	statusHandler := newStatusHandler(service.StatusService)
	damageHandler := newDamageHandler(service.DamageService)
	materialHandler := newMaterialHandler(service.MaterialService)
	statisticHandler := newStatisticHandler(service.StatisticService)
	historyHandler := newHistoryHandler(service.HistoryService)

	auth := router.Group("/auth")
	{
		auth.POST("/login", userHandler.LoginHandler)
		auth.POST("/register", userHandler.RegistrationHandler)
	}

	api := router.Group("/api", AuthMiddleware())
	{
		appeals := api.Group("/appeals")
		{
			appeals.POST("", appealHandler.CreateAppealHandler)
			appeals.GET("", appealHandler.GetAllAppealsHandler)
			appeals.GET("/:id", appealHandler.GetAppealHandler)
			appeals.PUT("", appealHandler.UpdateAppealHandler)
			appeals.DELETE("/:id", appealHandler.DeleteAppealHandler)
		}

		applications := api.Group("/applications")
		{
			applications.POST("", applicationHandler.CreateApplicationHandler)
			applications.GET("", applicationHandler.GetAllApplicationsHandler)
			applications.GET("/brigadir/:id_brigadir", applicationHandler.GetAllBrigadirApplicationsHandler)
			applications.GET("/:id", applicationHandler.GetApplicationHandler)
			applications.PUT("", applicationHandler.UpdateApplicationHandler)
			applications.DELETE("/:id", applicationHandler.DeleteApplicationHandler)
		}

		shutdowns := api.Group("/shutdowns")
		{
			shutdowns.POST("", shutdownHandler.CreateShutdownHandler)
			shutdowns.GET("", shutdownHandler.GetAllShutdownsHandler)
			shutdowns.GET("/:id", shutdownHandler.GetShutdownHandler)
			shutdowns.PUT("", shutdownHandler.UpdateShutdownHandler)
			shutdowns.DELETE("/:id", shutdownHandler.DeleteShutdownHandler)
		}

		accidents := api.Group("/accidents")
		{
			accidents.GET("", accidentHandler.GetAllAccidentsHandler)
		}

		characters := api.Group("/characters")
		{
			characters.GET("", characterHandler.GetAllCharactersHandler)
		}

		specializations := api.Group("/specializations")
		{
			specializations.GET("", specializationHandler.GetAllSpecializationsHandler)
		}

		importances := api.Group("/importances")
		{
			importances.GET("", importanceHandler.GetAllImportancesHandler)
		}

		statuses := api.Group("/statuses")
		{
			statuses.GET("", statusHandler.GetAllStatusHandler)
		}

		damages := api.Group("/damages")
		{
			damages.GET("", damageHandler.GetAllDamagesHandler)
		}

		materials := api.Group("/materials")
		{
			materials.GET("", materialHandler.GetAllMaterialsHandler)
		}

		statistics := api.Group("/statistics")
		{
			statistics.GET("", statisticHandler.GetStatisticHandler)
		}

		history := api.Group("/history")
		{
			history.GET("/appeals", historyHandler.GetAppeals)
			history.GET("/applications", historyHandler.GetApplications)
			history.GET("/shutdowns", historyHandler.GetShutdowns)
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
