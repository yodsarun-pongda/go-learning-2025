package service

import (
	"fmt"
	database "goapp/database"
	log "goapp/log"
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
	// c.IndentedJSON(http.StatusOK, users)

	rows, err := database.Db.Query("SELECT id, identify_id, ticket_id FROM customer")
	if err != nil {
		log.Error("Query error:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var identify_id, ticket_id string

		err := rows.Scan(&id, &identify_id, &ticket_id)
		if err != nil {
			log.Error("Row scan error:", err)
			continue
		}

		// Print ออก console
		fmt.Printf("ID: %d, identify_id: %s, ticket_id: %s\n", id, identify_id, ticket_id)
	}

	log.Info("Finished printing all customers")
}
