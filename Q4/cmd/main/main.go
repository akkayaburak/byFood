package main

import (
	service "byFood/Q4/pkg/services"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//adding cors
	router.Use(cors.Default())

	// Route handles & endpoints

	// Get all users
	router.GET("/users/", service.GetUsers)

	//Get a user
	router.GET("/users/:userid", service.GetUser)

	//Create a user
	router.POST("/users/", service.CreateUser)

	// Update a user
	router.PUT("/users/", service.UpdateUser)

	//Delete a specific user by the userID
	router.DELETE("/users/", service.DeleteUser)

	// serve the app
	fmt.Println("Server at 8080")
	router.Run("localhost:8080")

}
