package config

import (
	"strings"

	"github.com/gin-gonic/gin"
)

var Route *gin.Engine

func RouteConfig() {
	Route = gin.Default()
}

func StartRoute(portNumber string) {
	port := portNumber
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	Route.Run(port)
}
