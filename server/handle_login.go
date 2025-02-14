package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleLogin(c *gin.Context) {
	var user User
	query := "SELECT id, username, password FROM USERS WHERE username = $1"
	var id int
	var username string
	var password string
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}
	err = db.QueryRow(query, user.Username).Scan(&id, &username, &password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "username": username, "password": password})
}
