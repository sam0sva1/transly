package main

import (
	"net/http"
	"strconv"
	"transly/dbdrive"
	"transly/tools"

	"github.com/gin-gonic/gin"
)

func InitExerciseRoutes(server *Server) {
	router.GET("/exercises", func(c *gin.Context) {
		limitQuery := c.DefaultQuery("limit", "10")
		offsetQuery := c.DefaultQuery("offset", "-1")

		limit, err := strconv.Atoi(limitQuery)
		tools.Chk(err)
		offset, err := strconv.Atoi(offsetQuery)
		tools.Chk(err)

		params := dbdrive.ExerciseParams{
			Limit:  limit,
			Offset: offset,
		}
		exercises := server.exerciseService.GetCollection(params)

		c.JSON(http.StatusOK, gin.H{"exercises": exercises})
	})
}
