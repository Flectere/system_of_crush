package transport

import (
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type characterHandler struct {
	characterService *service.CharacterService
}

func newCharacterHandler(characterService *service.CharacterService) *characterHandler {
	return &characterHandler{characterService: characterService}
}

func (h *characterHandler) GetAllCharactersHandler(c *gin.Context) {
	characters, err := h.characterService.GetAllCharacters()
	if err != nil {
		c.JSON(400, err.Error())
	}

	c.JSON(200, characters)
}

func (h *characterHandler) GetCharacterHandler(c *gin.Context) {

}
