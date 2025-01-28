package transport

import (
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type historyHandler struct {
	historyService *service.HistoryService
}

func newHistoryHandler(historyService *service.HistoryService) *historyHandler {
	return &historyHandler{historyService: historyService}
}

func (h *historyHandler) GetAppeals(c *gin.Context) {
	appeals, err := h.historyService.GetAppeals()
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, appeals)
}

func (h *historyHandler) GetApplications(c *gin.Context) {
	applications, err := h.historyService.GetApplications()
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, applications)
}

func (h *historyHandler) GetShutdowns(c *gin.Context) {
	shutdowns, err := h.historyService.GetShutdowns()
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, shutdowns)
}
