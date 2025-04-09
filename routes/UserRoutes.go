package routes

import (
	"goapp/config"
	userService "goapp/service"

	log "github.com/sirupsen/logrus"
)

func InitUserRoutes() {
	getUsers()
	getUserByID()
}

func getUsers() {
	config.Route.GET("/users", userService.GetAllUsers)
	log.Info("GET /users route initialized")
}

func getUserByID() {
	config.Route.GET("/users/:id", userService.GetUserByID)
	log.Info("GET /users/:id route initialized")
}
