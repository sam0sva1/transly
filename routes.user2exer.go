package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUser2ExerRoutes(server *Server) {
	router.POST("/done", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Good boy!"})
	})
}
