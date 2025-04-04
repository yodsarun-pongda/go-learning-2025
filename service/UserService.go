package service

import (
	"goapp/model"
	"goapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []model.User{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	// Logic to get user by ID
	result := utils.Filter(users, func(user model.User) bool {
		return user.ID == id
	})
	c.IndentedJSON(http.StatusOK, result)
}

func GetAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
