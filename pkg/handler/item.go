package handler

import (
	"github.com/Serelllka/Federacion/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary AddItem
// @Security ApiKeyAuth
// @Tags item
// @Description adds item
// @ID addItem
// @Accept json
// @Produce json
// @Param input body entities.Item true "ItemId"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /items/ [post]
func (h *Handler) AddItem(c *gin.Context) {
	var input entities.Item

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Items.CreateItem(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary GetItem
// @Security ApiKeyAuth
// @Tags item
// @Description gets item
// @ID getItem
// @Produce json
// @Accept json
// @Success 200 {object} entities.Item
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Param id path int true "ItemId"
// @Router /items/{id} [get]
func (h *Handler) GetItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	item, err := h.services.Items.GetItemById(id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, item)
}

type getAllItemsResponse struct {
	Data []entities.Item `json:"data"`
}

// @Summary GetAllItems
// @Security ApiKeyAuth
// @Tags item
// @Description gets items
// @ID getAllItems
// @Produce json
// @Accept json
// @Success 200 {object} getAllItemsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /items/ [get]
func (h *Handler) GetAllItems(c *gin.Context) {
	items, err := h.services.Items.GetAllItems()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}
