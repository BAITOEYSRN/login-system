package main

import (
	"fmt"
	"log"
	"login-system/db"
	"login-system/middleware"
	"login-system/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(middleware.LoggerMiddleware())

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Database connected successfully")

	routes.Routes(r, db)

	r.Run(":8080")
}
