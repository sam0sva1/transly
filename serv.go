package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *Config
}

func (s *Server) Handler() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "It works!",
		})
	})

	return router
}

func (server *Server) Run() {
	fmt.Printf("Server starts on http://%s:%s", server.config.Host, server.config.Port)

	Engine := server.Handler()
	Engine.Run(server.config.Host + ":" + server.config.Port)
}
