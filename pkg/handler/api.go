package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type linkInput struct {
	Link string `json:"link" binding:"required"`
}

func (h *Handler) cropUrl(c *gin.Context) {
	var req linkInput

	if err := c.BindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	link, err := h.services.Link.CropLink(req.Link)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"url": link,
	})
}

func (h *Handler) getUrl(c *gin.Context) {
	var req linkInput

	if err := c.BindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	url, err := h.services.Link.GetLink(req.Link)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"url": url,
	})
}
