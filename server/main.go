package main

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Database connection setup
	connStr := "user=postgres password=4545 dbname=postgres sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	// Initialize Gin router
	r := gin.Default()
	// router config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	// Route to handle user registration
	r.POST("/register", registerUser)
	//router to handle user registration
	r.POST("/login", handleLogin)
	//router to handle user data update
	r.POST("/reset", handlePassword)
	// Start the server
	r.Run(":8080")
}

// Struct to receive JSON request
