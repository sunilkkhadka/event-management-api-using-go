package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunilkkhadka/gin-practice/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run()
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "resource does not exist",
		})
		return
	}

	event.ID = 1
	event.UserID = 1
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event Created Successfully",
		"event":   event,
	})
}
