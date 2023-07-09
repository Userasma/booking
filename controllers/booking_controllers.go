package controllers

import (
	"github.com/gin-gonic/gin"
)

// Handler for the home page
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// Handler for creating a new booking
func CreateBookingHandler(c *gin.Context) {
	// Retrieve booking details from the request
	// Process and save the booking in the database

	// Return a response
	c.JSON(http.StatusOK, gin.H{
		"message": "Booking created successfully",
	})
}
