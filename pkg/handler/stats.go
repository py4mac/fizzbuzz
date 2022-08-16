package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Stats(c *gin.Context) {
	stats, err := h.statsRepo.Process()
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, stats)
}
