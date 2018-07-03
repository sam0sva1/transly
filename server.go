package main

import (
	"fmt"
	"net/http"
	"strconv"

	"transly/config"
	"transly/dbdrive"
	"transly/tools"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config          *config.Config
	exerciseService *dbdrive.ExerciseService
}

func (s *Server) Handler() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "It works!",
		})
	})

	router.GET("/exercises", func(c *gin.Context) {
		limitQuery := c.DefaultQuery("limit", "10")
		offsetQuery := c.DefaultQuery("offset", "0")
		// fmt.Println("limit and offset", limit, offset)
		limit, err := strconv.Atoi(limitQuery)
		tools.Chk(err)
		offset, err := strconv.Atoi(offsetQuery)
		tools.Chk(err)

		exercises := s.exerciseService.GetCollection(limit, offset)

		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

		c.JSON(http.StatusOK, gin.H{"exercises": exercises})
	})

	return router
}

func (server *Server) Run() {
	fmt.Printf("Server starts on http://%s:%s", server.config.Host, server.config.Port)

	Engine := server.Handler()

	Engine.Run(server.config.Host + ":" + server.config.Port)
}
