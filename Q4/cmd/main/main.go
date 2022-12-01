package main

import (
	service "byFood/Q4/pkg/services"
	"fmt"

	"github.com/gin-gonic/gin"
	//_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	// Route handles & endpoints

	// Get all users
	router.GET("/users/", service.GetUsers)

	//Create a user
	router.POST("/users/", service.CreateUser)

	// Update a user
	router.PUT("/users/", service.UpdateUser)

	// Delete a specific user by the userID
	//router.DELETE("/movies/{movieid}", "")

	// serve the app
	fmt.Println("Server at 8080")
	router.Run("localhost:8080")

}
