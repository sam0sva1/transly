package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"transly/config"
	"transly/dbdrive"
	"transly/tools"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config          *config.Config
	exerciseService *dbdrive.ExerciseService
	userService     *dbdrive.UserService
}

func (s *Server) Handlers() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "It works!",
		})
	})

	router.GET("/exercises", func(c *gin.Context) {
		limitQuery := c.DefaultQuery("limit", "10")
		offsetQuery := c.DefaultQuery("offset", "-1")

		// fmt.Println("limit and offset", limit, offset)
		limit, err := strconv.Atoi(limitQuery)
		tools.Chk(err)
		offset, err := strconv.Atoi(offsetQuery)
		tools.Chk(err)

		params := dbdrive.ExerciseParams{
			Limit:  limit,
			Offset: offset,
		}
		exercises := s.exerciseService.GetCollection(params)

		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

		c.JSON(http.StatusOK, gin.H{"exercises": exercises})
	})

	router.GET("/user/:id", func(c *gin.Context) {
		strId := c.Param("id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			// tools.Chk(err)
			c.String(http.StatusNotFound, "Invalid id.")
			return
		}

		user, ok := s.userService.GetUserById(id)
		if ok {
			c.String(http.StatusNotFound, "No user with this id")
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	router.POST("/user", func(c *gin.Context) {
		var user *dbdrive.User
		if err := c.BindJSON(&user); err == nil {
			var validations []string

			if user.Login == "" {
				validations = append(validations, "login")
			}
			if user.Password == "" {
				validations = append(validations, "password")
			}

			if len(validations) > 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong fields: " + strings.Join(validations, ", ")})
				return
			}

			isExist := s.userService.Check(user)
			if isExist {
				c.JSON(http.StatusBadRequest, gin.H{"error": "This login is in use."})
				return
			}

			err = s.userService.Create(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			c.JSON(http.StatusOK, gin.H{"user": user})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	return router
}

func (server *Server) Run() {
	fmt.Printf("Server starts on http://%s:%s", server.config.Host, server.config.Port)

	Engine := server.Handlers()

	Engine.Run(server.config.Host + ":" + server.config.Port)
}
