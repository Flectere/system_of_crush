package transport

import (
	"net/http"

	"github.com/Flectere/system_of_crush/internal/models"
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type brigadeHandler struct {
	brigadeService *service.BrigadeService
}

func newBrigadeHandler(brigadeService *service.BrigadeService) *brigadeHandler {
	return &brigadeHandler{brigadeService: brigadeService}
}

func (h *brigadeHandler) GetAllBrigadesHandler(c *gin.Context) {
	brigades, err := h.brigadeService.GetAllBrigades()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, brigades)
}

func (h *brigadeHandler) GetBrigadeByIDHandler(c *gin.Context) {
	id := c.Param("id")
	brigade, err := h.brigadeService.GetBrigadeByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, brigade)
}

func (h *brigadeHandler) CreateBrigadeHandler(c *gin.Context) {
	var newBrigade models.BrigadeCreate
	c.ShouldBindJSON(&newBrigade)

	id, err := h.brigadeService.CreateBrigade(newBrigade)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(201, gin.H{"brigade_id": id})
}

func (h *brigadeHandler) EditBrigadeHandler(c *gin.Context) {
	var brigade models.BrigadeCreate
	c.ShouldBindJSON(&brigade)

	id := c.Param("id")

	err := h.brigadeService.EditBrigade(id, brigade)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"brigade_id": id})
}

func (h *brigadeHandler) GetFreeBrigadirsHandler(c *gin.Context) {
	var brigadirs []models.Brigadir

	brigadirs, err := h.brigadeService.GetFreeBrigadirs()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}

	c.JSON(200, brigadirs)
}
