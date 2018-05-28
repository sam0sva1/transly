package main

import (
	"fmt"
	"net/http"
	"transly/config"
	"transly/dbdrive"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config          *config.Config
	exerciseService *dbdrive.ExerciseService
}

func (s *Server) Handler() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "It works!",
		})
	})

	router.GET("/exercises", func(c *gin.Context) {
		exercises := s.exerciseService.GetCollection()
		c.JSON(http.StatusOK, gin.H{"exercises": exercises})
	})

	return router
}

func (server *Server) Run() {
	fmt.Printf("Server starts on http://%s:%s", server.config.Host, server.config.Port)

	Engine := server.Handler()
	Engine.Run(server.config.Host + ":" + server.config.Port)
}
