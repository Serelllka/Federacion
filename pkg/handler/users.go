package handler

import (
	federacion "github.com/Serelllka/Federacion/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAllUsersResponse struct {
	Data []federacion.UserInfo `json:"data"`
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.Users.GetAllUsers()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllUsersResponse{
		Data: users,
	})
}
