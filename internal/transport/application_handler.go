package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Flectere/system_of_crush/internal/models"
	"github.com/Flectere/system_of_crush/internal/service"
)

type applicationHandler struct {
	applicationService *service.ApplicationService
}

func newApplicationHandler(applicationService *service.ApplicationService) *applicationHandler {
	return &applicationHandler{applicationService: applicationService}
}

// Обработчик для обработки запросов на создание новой заявки
func (h *applicationHandler) CreateApplicationHandler(c *gin.Context) {
	var application models.ApplicationCreate

	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	applicationID, err := h.applicationService.CreateApplication(application)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"application_id": applicationID})
}

// Обработчик для обработки запросов на получение всех заявок
func (h *applicationHandler) GetAllApplicationsHandler(c *gin.Context) {
	applications, err := h.applicationService.GetAllApplications()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, applications)
}

// Обработчик для обработки запросов на получение всех заявок у бригадира
func (h *applicationHandler) GetAllBrigadirApplicationsHandler(c *gin.Context) {
	brigadirID := c.Param("id_brigadir")
	applications, err := h.applicationService.GetAllBrigadirApplications(brigadirID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
	}

	c.JSON(200, applications)
}

// Обработчик для получения всех свободных заявок
func (h *applicationHandler) GetFreeApplicationsHandler(c *gin.Context) {
	applications, err := h.applicationService.GetAllFreeApplications()
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, applications)
}

// Обработчик для обработки запросов получения конкретной заявки
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
	var application models.ApplicationEdit

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

func (h *applicationHandler) SetApplicationToBrigadirHandler(c *gin.Context) {
	applicationID := c.Param("id")

	type requestBody struct {
		BrigadirID string `json:"id_brigadir"`
	}
	var body requestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := h.applicationService.SetApplicationToBrigadir(body.BrigadirID, applicationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
