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
