package app

import (
	"fmt"
	"log"

	"github.com/Flectere/system_of_crush/config"
	"github.com/Flectere/system_of_crush/internal/database"
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/Flectere/system_of_crush/internal/transport"
)

func Run(configPath string) {
	err := config.InitConfig(configPath)
	if err != nil {
		log.Fatal("Ошибка при загрузке конфигурации: ", err)
	}
	databaseConfig := config.Config.DataBaseConfig
	serverConfig := config.Config.ServerConfig

	db := database.InitDB(databaseConfig)
	defer database.CloseDB()

	service := service.NewService(db)

	router := transport.NewRouter(service)

	connString := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)

	err = router.Run(connString)
	if err != nil {
		log.Fatal("Ошибка при запуске сервера ", err)
	}
}
