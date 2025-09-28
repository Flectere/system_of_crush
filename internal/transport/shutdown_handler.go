package transport

import (
	"github.com/Flectere/system_of_crush/internal/models"
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type shutdownHandler struct {
	shutdownService *service.ShutdownService
}

func newShutdownHandler(shutdownService *service.ShutdownService) *shutdownHandler {
	return &shutdownHandler{shutdownService: shutdownService}
}

// Обработчик для обработки запросов на создание отключения
func (h *shutdownHandler) CreateShutdownHandler(c *gin.Context) {
	var Shutdown models.ShutdownCreate

	if err := c.ShouldBindJSON(&Shutdown); err != nil {
		c.JSON(400, err.Error())
		return
	}

	ShutdownID, err := h.shutdownService.CreateShutdown(Shutdown)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"Shutdown_id": ShutdownID})
}

// Обработчик для обработки запросов на получение всех отключений
func (h *shutdownHandler) GetAllShutdownsHandler(c *gin.Context) {
	shutdowns, err := h.shutdownService.GetAllShutdowns()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, shutdowns)
}

// TODO подумать над атомарными данными
// Обработчик для обработки запросов получения конкретного отключения
func (h *shutdownHandler) GetShutdownHandler(c *gin.Context) {
	id := c.Param("id")

	Shutdown, err := h.shutdownService.GetShutdownById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Shutdown)
}

// TODO подумать про атомарность
func (h *shutdownHandler) UpdateShutdownHandler(c *gin.Context) {
	var Shutdown models.Shutdown

	if err := c.ShouldBindJSON(&Shutdown); err != nil {
		c.JSON(400, gin.H{"error": "Неверные данные"})
		return
	}

	err := h.shutdownService.UpdateShutdown(Shutdown)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Shutdown)
}

func (h *shutdownHandler) DeleteShutdownHandler(c *gin.Context) {
	id := c.Param("id")

	err := h.shutdownService.DeleteShutdown(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}
