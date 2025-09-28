package transport

import (
	"net/http"

	"github.com/Flectere/system_of_crush/internal/models"
	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

type mobileHandler struct {
	applicationService *service.ApplicationService
}

func newMobileHandler(applicationService *service.ApplicationService) *mobileHandler {
	return &mobileHandler{applicationService: applicationService}
}

func (h *mobileHandler) GetAllBrigadirApplicationsHandler(c *gin.Context) {
	brigadirID := c.Param("id_brigadir")
	applications, err := h.applicationService.GetAllBrigadirApplications(brigadirID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
	}

	c.JSON(200, applications)
}

func (h *mobileHandler) SetApplicationToBrigadir(c *gin.Context) {
	applicationID := c.Param("id")

	type requestBody struct {
		BrigadirID string `json:"id_brigadir"`
	}
	var body requestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := h.applicationService.SetApplicationToBrigadir(body.BrigadirID, applicationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *mobileHandler) GetFreeApplicationsHandler(c *gin.Context) {
	applications, err := h.applicationService.GetAllFreeApplications()
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(200, applications)
}

func (h *mobileHandler) GetApplicationHandler(c *gin.Context) {
	var application models.ApplicationBrigadir

	id := c.Param("id")
	application, err := h.applicationService.GetApllicationByIdMobile(id)

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, application)
}

func (h *mobileHandler) StartApplicationHandler(c *gin.Context) {
	id := c.Param("id")

	err := h.applicationService.StartApplication(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(204)
}

func (h *mobileHandler) FinishApplicationHandler(c *gin.Context) {
	id := c.Param("id")

	err := h.applicationService.FinishApplication(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(204)
}
