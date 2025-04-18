package main

import (
	"goapp/first_program"
	log "goapp/log"
	"goapp/routes"
	"goapp/service"
	"goapp/utils"

	config "goapp/config"
)

func main() {
	// Initialize the configuration
	log.InitLogConfig()

	// Hello world program
	tutorial_1()

	// TODO: Connect to DB
	// database.MysqlConfiguration()
	// defer database.Db.Close()

	// TODO goroutine
	threadNumber := 20
	service.PocMultiThread(threadNumber)

	// TODO Garbage

	// TODO channel
	// TODO select
	// TODO sync.Mutex
	// TODO sync.WaitGroup

	log.Info("This is an info message")

	// inject Rest API
	// startRoutes()
}

func startRoutes() {
	// Initialize the routes
	config.RouteConfig()

	// Initialize the routes
	routes.InitUserRoutes()

	log.Info("All Routes initialized")
	config.StartRoute("8080")
}

func tutorial_1() {
	log.Info("MAIN said:: Hello, World!")
	log.Info(first_program.TestPrint())

	result := utils.Filter([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i%2 == 0
	})
	log.Info(result)
}
