package transport

import (
	"github.com/gin-gonic/gin"

	"github.com/Flectere/system_of_crush/internal/models"
	"github.com/Flectere/system_of_crush/internal/service"
)

type applicationHandler struct {
	applicationService *service.ApplicationService
}

func NewApplicationHandler(applicationService *service.ApplicationService) *applicationHandler {
	return &applicationHandler{applicationService: applicationService}
}

// Обработчик для обработки запросов на создание нового обращения
func (h *applicationHandler) CreateApplicationHandler(c *gin.Context) {
	var application models.Application

	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(400, gin.H{"error": "Неверные данные"})
		return
	}

	applicationID, err := h.applicationService.CreateApplication(application)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"application_id": applicationID})
}

// Обработчик для обработки запросов на получение всех обращений
func (h *applicationHandler) GetAllApplicationsHandler(c *gin.Context) {

	applications, err := h.applicationService.GetAllApplications()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, applications)
}

// Обработчик для обработки запросов получения конкретного обращения
func (h *applicationHandler) GetApplicationHandler(c *gin.Context) {
	var application models.Application

	id := c.Param("id")
	application, err := h.applicationService.GetApplicationById(id)

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, application)
}

func (h *applicationHandler) UpdateApplicationHandler(c *gin.Context) {
	var application models.Application

	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(400, gin.H{"error": "Неверные данные"})
		return
	}

	err := h.applicationService.UpdateApplication(application)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, application)
}

func (h *applicationHandler) DeleteApplicationHandler(c *gin.Context) {
	id := c.Param("id")

	err := h.applicationService.DeleteApplication(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}
