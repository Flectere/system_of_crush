package transport

import (
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type statisticHandler struct {
	statisticService *service.StatisticService
}

func newStatisticHandler(statisticService *service.StatisticService) *statisticHandler {
	return &statisticHandler{statisticService: statisticService}
}

func (h *statisticHandler) GetStatisticHandler(c *gin.Context) {
	statistic := h.statisticService.GetStatistics()

	c.JSON(200, statistic)
}
