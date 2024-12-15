package transport

import (
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type specializationHandler struct {
	specializationService *service.SpecializationService
}

func newSpecializationHandler(specializationService *service.SpecializationService) *specializationHandler {
	return &specializationHandler{specializationService: specializationService}
}

func (h *specializationHandler) GetAllSpecializationsHandler(c *gin.Context) {
	specializations, err := h.specializationService.GetAllSpecializations()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, specializations)

}

func (h *specializationHandler) GetSpecializationHandler(c *gin.Context) {

}
