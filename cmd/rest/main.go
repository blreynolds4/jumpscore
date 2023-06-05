package main

import (
	"fmt"
	"jumpscore/jumpscore"
	"jumpscore/store"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// methods to handle rest requests
func createEvent(c *gin.Context) {
	fmt.Println("Handling get events")
	c.Header("Content-Type", "application/json")
	var e jumpscore.Event
	err := c.BindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	err = store.CreateEvent(e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, e)
}
func getEvents(c *gin.Context) {
	fmt.Println("Handling get events")
	store.GetEvents()
}

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	api.GET("/events", getEvents)
	api.POST("/events", createEvent)

	router.Run("localhost:8080")
}
