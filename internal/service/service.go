package service

type Service struct {
	UserService           *UserService
	AppealService         *AppealService
	ApplicationService    *ApplicationService
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
