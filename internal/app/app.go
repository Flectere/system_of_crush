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
	applicationService := service.NewApplicationService(db)
	shutdownService := service.NewShutdownService(db)
	accidentService := service.NewAccidentService(db)
	characterService := service.NewCharacterService(db)
	specializationService := service.NewSpecializationService(db)
	importanceService := service.NewImportanceService(db)
	statusService := service.NewStatusService(db)
	materialService := service.NewMaterialService(db)
	damageService := service.NewDamageService(db)
	statisticService := service.NewStatisticService(db)
	historyService := service.NewHistoryService(db)

	service := service.Service{
		UserService:           userService,
		AppealService:         appealService,
		ApplicationService:    applicationService,
		ShutdownService:       shutdownService,
		AccidentService:       accidentService,
		CharacterService:      characterService,
		SpecializationService: specializationService,
		ImportanceService:     importanceService,
		StatusService:         statusService,
		MaterialService:       materialService,
		DamageService:         damageService,
		StatisticService:      statisticService,
		HistoryService:        historyService,
	}

	router := transport.NewRouter(&service)

	err = router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal("Ошибка при запуске сервера ", err)
	}
}
