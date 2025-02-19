package transport

import (
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type accidentHandler struct {
	accidentService *service.AccidentService
}

func newAccidentHandler(accidentService *service.AccidentService) *accidentHandler {
	return &accidentHandler{accidentService: accidentService}
}

func (h *accidentHandler) GetAllAccidentsHandler(c *gin.Context) {
	accidents, err := h.accidentService.GetAllAccidents()

	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, accidents)
}
