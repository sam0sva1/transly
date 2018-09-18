package serving

import (
	"fmt"
	"net/http"

	"transly/config"
	"transly/dbdrive"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Config          *config.Config
	ExerciseService *dbdrive.ExerciseService
	UserService     *dbdrive.UserService
}

var router *gin.Engine

func (s *Server) Handlers() *gin.Engine {
	router = gin.Default()

	router.Use(cors.Default())

	InitUserRoutes(s)
	InitExerciseRoutes(s)
	InitUser2ExerRoutes(s)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "It works!",
		})
	})

	return router
}

func (server *Server) Run() {
	fmt.Printf("Server starts on http://%s:%s", server.Config.Host, server.Config.Port)

	Engine := server.Handlers()

	Engine.Run(server.Config.Host + ":" + server.Config.Port)
}
