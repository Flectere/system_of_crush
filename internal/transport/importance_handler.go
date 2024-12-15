package transport

import (
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type importanceHandler struct {
	importanceService *service.ImportanceService
}

func newImportanceHandler(importanceService *service.ImportanceService) *importanceHandler {
	return &importanceHandler{importanceService: importanceService}
}

func (h *importanceHandler) GetAllImportancesHandler(c *gin.Context) {
	importances, err := h.importanceService.GetAllImportances()
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, importances)
}

func (h *importanceHandler) GetImportanceHandler(c *gin.Context) {

}
