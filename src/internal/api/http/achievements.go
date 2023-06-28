package http

import (
	"github.com/Serelllka/Federacion/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAllAchievementsResponse struct {
	Data []entities.Achievement `json:"data"`
}

func (h *Handler) GetAchievements(c *gin.Context) {
	achievements, err := h.services.GetAllAchievements()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllAchievementsResponse{
		Data: achievements,
	})
}

func (h *Handler) GetAchievement(c *gin.Context) {

}

func (h *Handler) AddAchievement(c *gin.Context) {

}
