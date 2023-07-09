package main

import (
	"booking-system/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Routes
	router.GET("/", controllers.IndexHandler)
	router.POST("/booking", controllers.CreateBookingHandler)

	router.Run(":8080")
}
