package main

import (
	"fmt"
	"jumpscore/jumpscore"
	"jumpscore/store"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func getEventStore() store.EventStore {
	store, err := store.NewFileEventStore("zEventStore")
	if err != nil {
		panic(err)
	}

	return store
}

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

	eventStore := getEventStore()
	err = eventStore.CreateEvent(e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, e)
}
func getEvents(c *gin.Context) {
	fmt.Println("Handling get events")
	eventStore := getEventStore()
	events, err := eventStore.GetEvents()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	// need a response object maybe to hold the array?
	// don't like to return bare arrays
	c.JSON(http.StatusOK, events)
}

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./jumpscore_ui", true)))

	// Setup route group for the API
	api := router.Group("/api")
	api.Use(cors.Default())
	api.GET("/events", getEvents)
	api.POST("/events", createEvent)

	router.Run("localhost:8080")
}
