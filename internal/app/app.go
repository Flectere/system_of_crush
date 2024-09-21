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

	db := database.InitDB(databaseConfig)
	userService := service.NewUserService(db)
	router := transport.NewRouter(userService)

	err = router.Run("25.8.104.246:8080")
	if err != nil {
		log.Fatal("Ошибка при запуске сервера", err)
	}
}
