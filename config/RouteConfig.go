package config

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func RouteConfig() *gin.Engine {
	return gin.Default()
}

func StartRoute(route *gin.Engine, portNumber string) {
	port := portNumber
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	route.Run(port)
}
