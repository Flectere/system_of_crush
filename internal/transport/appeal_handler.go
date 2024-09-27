package transport

import (
	"github.com/Flectere/system_of_crush/internal/models"
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type appealHandler struct {
	appealService *service.AppealService
}

func NewAppealHandler(appealService *service.AppealService) *appealHandler {
	return &appealHandler{appealService: appealService}
}

// Обработчик для обработки запросов на создание нового обращения
func (h *appealHandler) CreateAppealHandler(c *gin.Context) {
	// Получение данных обращения из запроса
	var appeal models.Appeal
	if err := c.ShouldBindJSON(&appeal); err != nil {
		c.JSON(400, gin.H{"error": "Неверные данные"})
		return
	}

	// Проверяем, что ID специализации и важности присутствуют
	if appeal.Specialization.ID == 0 || appeal.Importance.ID == 0 {
		c.JSON(400, gin.H{"error": "Необходимо указать специализацию и важность"})
		return
	}

	// Создание нового обращения
	appealID, err := h.appealService.CreateAppeal(appeal)
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка при создании обращения"})
		return
	}

	c.JSON(201, gin.H{"appeal_id": appealID})
}

// Обработчик для обработки запросов на получение всех обращений
func (h *appealHandler) GetAllAppealsHandler(c *gin.Context) {
	appeals, err := h.appealService.GetAllAppeals()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, appeals)
}

// Обработчик для обработки запросов получения конкретного обращения
func (h *appealHandler) GetAppealHandler(c *gin.Context) {
	var appeal models.Appeal

	id := c.Param("id")
	appeal, err := h.appealService.GetAppealById(id)

	if err != nil {
		c.JSON(404, gin.H{"error": "Обращение не найдено"})
		return
	}

	c.JSON(200, appeal)

}

func (h *appealHandler) UpdateAppealHandler(c *gin.Context) {

}

func (h *appealHandler) DeleteAppealHandler(c *gin.Context) {

}
