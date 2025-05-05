package main

import (
	"goapp/first_program"
	log "goapp/log_config"
	"goapp/routes"
	"goapp/service"
	_ "goapp/service"
	"goapp/utils"

	config "goapp/config"
)

func main() {
	// Initialize the configuration
	log.InitLogConfig()

	// Hello world program
	// tutorial_1()

	// ===========================
	// Basic Array
	// log.Info("Basic ==========================")
	// names := []string{"John", "Doe", "Smith"}
	// names = append(names, "dd")
	// names = append(names, "ddd")
	// names = append(names, "Smfdffith")
	// customerListsJson, _ := json.Marshal(names)
	// log.Info(string(customerListsJson))

	// // Basic Map
	// mapNames := make(map[string]int) // map[key]value
	// mapNames["John"] = 1
	// mapNames["Doe"] = 2
	// mapNames["Smith"] = 3
	// log.Info(mapNames)
	// delete(mapNames, "Smith")
	// log.Info(mapNames)

	// log.Info("==========================")
	// ===========================

	// Connect to DB
	// database.MysqlConfiguration()
	// defer database.Db.Close()

	// goroutine
	// threadNumber := 20
	// service.PocMultiThread(threadNumber)

	// TODO chanel
	service.PocChannel()
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
