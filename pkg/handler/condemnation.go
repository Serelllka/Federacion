package handler

import (
	"github.com/Serelllka/Federacion/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) AddCondemnation(c *gin.Context) {
	var input entities.Condemnation

	condemnerId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.Time = time.Now()
	input.CondemnerId = condemnerId

	id, err := h.services.Condemnations.CreateCondemnation(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type GetAllCondemnationsResponse struct {
	Data []entities.Condemnation `json:"data"`
}

// @Summary GetAllCondemnations
// @Tags condemn
// @Description GetAllCondemnations
// @ID getAllCondemns
// @Accept json
// @Produce json
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) GetAllCondemnations(c *gin.Context) {
	condemnations, err := h.services.Condemnations.GetAllCondemnations()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetAllCondemnationsResponse{
		Data: condemnations,
	})
}

func (h *Handler) GetCondemnation(c *gin.Context) {
}
