package main

import (
	common "byFood/Q4/pkg/common"
	persistence "byFood/Q4/pkg/persistence"
	service "byFood/Q4/pkg/services"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db := persistence.SetupDB()

	_, err := db.Query(`CREATE TABLE IF NOT EXISTS public.user
	(
		user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
		username text COLLATE pg_catalog."default" NOT NULL,
		password_hash text COLLATE pg_catalog."default" NOT NULL,
		mail text COLLATE pg_catalog."default" NOT NULL,
		is_deleted boolean NOT NULL DEFAULT false,
		CONSTRAINT user_id_pkey PRIMARY KEY (user_id)
	)
	
	TABLESPACE pg_default;
	
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	`)

	common.CheckError(err)

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
