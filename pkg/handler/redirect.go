package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) redirectLink(c *gin.Context) {
	code := c.Param("code")

	redirectUrl, err := h.services.Link.RedirectLink(code)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid url code")
	}

	c.Redirect(301, redirectUrl)
}
