package main

import (
	"goapp/first_program"
	"goapp/routes"
	"goapp/utils"

	config "goapp/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Initialize the configuration
	config.InitLogConfig()

	// Hello world program
	tutorial_1()

	// inject Rest API
	initRoutes()

	// TODO goroutine
	// TODO Garbage
	// TODO channel
	// TODO select
	// TODO sync.Mutex
	// TODO sync.WaitGroup

	log.Info("This is an info message")
}

func initRoutes() {
	// Initialize the routes
	engine := config.RouteConfig()

	// Initialize the routes
	routes.InitUserRoutes(engine)

	log.Info("All Routes initialized")
	config.StartRoute(engine, "8080")
}

func tutorial_1() {
	log.Info("MAIN said:: Hello, World!")
	log.Info(first_program.TestPrint())

	result := utils.Filter([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i%2 == 0
	})
	log.Info(result)
}
