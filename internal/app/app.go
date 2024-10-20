package app

import (
	"log"

	"github.com/Flectere/system_of_crush/configs"
	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/Flectere/system_of_crush/internal/transport"
)

func Run(configPath string) {

	config, err := configs.LoadConfig(configPath)
	if err != nil {
		log.Fatal("Ошибка при загрузке конфигурации: ", err)
	}

	db := database.InitDB(config)
	defer database.CloseDB()

	userService := service.NewUserService(db)
	appealService := service.NewAppealService(db)

	service := service.Service{
		UserService:   userService,
		AppealService: appealService,
	}

	router := transport.NewRouter(&service)

	err = router.Run("25.19.79.114:8080")
	if err != nil {
		log.Fatal("Ошибка при запуске сервера ", err)
	}
}
