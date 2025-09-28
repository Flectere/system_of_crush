package transport

import (
	"github.com/Flectere/system_of_crush/internal/models"
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type appealHandler struct {
	appealService *service.AppealService
}

func newAppealHandler(appealService *service.AppealService) *appealHandler {
	return &appealHandler{appealService: appealService}
}

// Обработчик для обработки запросов на создание нового обращения
func (h *appealHandler) CreateAppealHandler(c *gin.Context) {
	var appeal models.AppealCreate

	if err := c.ShouldBindJSON(&appeal); err != nil {
		c.JSON(400, err.Error())
		return
	}

	appealID, err := h.appealService.CreateAppeal(appeal)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
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
	id := c.Param("id")

	appeal, err := h.appealService.GetAppealById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, appeal)
}

func (h *appealHandler) UpdateAppealHandler(c *gin.Context) {
	var appeal models.AppealEdit

	if err := c.ShouldBindJSON(&appeal); err != nil {
		c.JSON(400, gin.H{"error": "Неверные данные"})
		return
	}

	err := h.appealService.UpdateAppeal(appeal)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, appeal)
}

func (h *appealHandler) DeleteAppealHandler(c *gin.Context) {
	id := c.Param("id")

	err := h.appealService.DeleteAppeal(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}
