package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func registerUser(c *gin.Context) {
	var user User

	// Bind JSON request to struct
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Insert into the database
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	var id int
	err = db.QueryRow(query, user.Username, string(hashedPassword)).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user_id": id})
}
