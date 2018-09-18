package serving

import (
	"net/http"
	"strconv"
	"strings"
	"transly/dbdrive"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(server *Server) {
	router.GET("/user/:id", func(c *gin.Context) {
		strId := c.Param("id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			// tools.Chk(err)
			c.String(http.StatusNotFound, "Invalid id.")
			return
		}

		user, ok := server.UserService.GetUserById(id)
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

			isExist := server.UserService.Check(user)
			if isExist {
				c.JSON(http.StatusBadRequest, gin.H{"error": "This login is in use."})
				return
			}

			err = server.UserService.Create(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			c.JSON(http.StatusOK, gin.H{"user": user})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
}
