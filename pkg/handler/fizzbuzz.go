package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/py4mac/fizzbuzz/pkg/fizzbuzz"
)

func (h *Handler) Fizzbuz(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(h.timeoutMs)*time.Millisecond)
	defer cancel()

	var req fizzbuzz.Fizzbuz
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	ret, err := req.Process(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := h.statsRepo.Record(req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, ret)
}
