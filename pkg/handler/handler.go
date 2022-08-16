package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/py4mac/fizzbuzz/pkg/stats"
)

type Handler struct {
	Engine    *gin.Engine
	statsRepo stats.Stats
	timeoutMs int32
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func NewHandler(statsRepo stats.Stats, timeoutMs int32) *Handler {
	h := &Handler{}
	e := gin.Default()
	e.Use(gin.Recovery())
	e.GET("/health", Health)

	g := e.Group("/api/v1")
	g.GET("/fizzbuzz", h.Fizzbuz)
	g.GET("/stats", h.Stats)

	h.Engine = e
	h.statsRepo = statsRepo
	h.timeoutMs = timeoutMs

	return h
}
