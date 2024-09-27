package app

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/Flectere/system_of_crush/internal/transport"

	"github.com/Flectere/system_of_crush/internal/database"
)

func Run(configPath string) {

	var databaseConfig database.DataBaseConfig

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatal("Не удалось открыть конфиг", err)
	}
	defer configFile.Close()

	configData, err := io.ReadAll(configFile)
	if err != nil {
		log.Fatal("Не удалось прочитать данные из конфига", err)
	}

	err = json.Unmarshal(configData, &databaseConfig)
	if err != nil {
		log.Fatal("Не удалось десериализовать конфиг", err)
	}
	// Инициализация базы данных
	db := database.InitDB(databaseConfig)
	defer database.CloseDB()

	// Инициализация сервисов
	userService := service.NewUserService(db)
	appealService := service.NewAppealService(db)

	service := service.Service{
		UserService:   userService,
		AppealService: appealService,
	}

	router := transport.NewRouter(&service)

	err = router.Run("25.19.79.114:8080")
	if err != nil {
		log.Fatal("Ошибка при запуске сервера", err)
	}
}
