package transport

import (
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type damageHandler struct {
	damageService *service.DamageService
}

func newDamageHandler(damageService *service.DamageService) *damageHandler {
	return &damageHandler{damageService: damageService}
}

func (h *damageHandler) GetAllDamagesHandler(c *gin.Context) {
	characters, err := h.damageService.GetAllDamages()
	if err != nil {
		c.JSON(400, err.Error())
	}

	c.JSON(200, characters)
}
