package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddFriend(c *gin.Context) {
	id, _ := c.Get(userContext)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
