package routes

import (
	userService "goapp/service"

	// "net/http"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InitUserRoutes(route *gin.Engine) {
	getUsers(route)
	getUserByID(route)
}

func getUsers(route *gin.Engine) {
	route.GET("/users", userService.GetAllUsers)
	log.Info("GET /users route initialized")
}

func getUserByID(route *gin.Engine) {
	route.GET("/users/:id", userService.GetUserByID)
	log.Info("GET /users/:id route initialized")
}
