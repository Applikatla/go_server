package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlePassword(c *gin.Context) {
	var user User
	query := "SELECT USERNAME FROM USERS WHERE ID = $1"
	var username string
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Database error"})
		return
	}
	err = db.QueryRow(query, user.Id).Scan(&username)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusFound, gin.H{"message": "Resent link will send shortly", "user name": username})
}
