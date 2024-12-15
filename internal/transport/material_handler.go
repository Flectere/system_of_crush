package transport

import (
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type materialHandler struct {
	materialService *service.MaterialService
}

func newMaterialHandler(materialService *service.MaterialService) *materialHandler {
	return &materialHandler{materialService: materialService}
}

func (h *materialHandler) GetAllMaterialsHandler(c *gin.Context) {
	materials, err := h.materialService.GetAllMaterials()
	if err != nil {
		c.JSON(400, err.Error())
	}

	c.JSON(200, materials)
}
