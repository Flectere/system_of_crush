package service

import (
	"github.com/Flectere/system_of_crush/internal/database"
)

type Service struct {
	UserService           *UserService
	AppealService         *AppealService
	ApplicationService    *ApplicationService
	BrigadeService        *BrigadeService
	ShutdownService       *ShutdownService
	SpecializationService *SpecializationService
	ImportanceService     *ImportanceService
	AccidentService       *AccidentService
	StatusService         *StatusService
	CharacterService      *CharacterService
	DamageService         *DamageService
	MaterialService       *MaterialService
	StatisticService      *StatisticService
	HistoryService        *HistoryService
}

func NewService(db *database.Database) *Service {
	return &Service{
		UserService:           NewUserService(db),
		AppealService:         NewAppealService(db),
		ApplicationService:    NewApplicationService(db),
		BrigadeService:        NewBrigadeService(db),
		ShutdownService:       NewShutdownService(db),
		AccidentService:       NewAccidentService(db),
		CharacterService:      NewCharacterService(db),
		SpecializationService: NewSpecializationService(db),
		ImportanceService:     NewImportanceService(db),
		StatusService:         NewStatusService(db),
		MaterialService:       NewMaterialService(db),
		DamageService:         NewDamageService(db),
		StatisticService:      NewStatisticService(db),
		HistoryService:        NewHistoryService(db),
	}
}
