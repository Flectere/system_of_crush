package transport

import (
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type statusHandler struct {
	statusService *service.StatusService
}

func newStatusHandler(statusService *service.StatusService) *statusHandler {
	return &statusHandler{statusService: statusService}
}

func (h *statusHandler) GetAllStatusHandler(c *gin.Context) {
	statuses, err := h.statusService.GetAllStatuses()
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, statuses)
}

func (h *statusHandler) GetStatusHandler(c *gin.Context) {

}
