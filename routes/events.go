package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jarqvi/gin-crud/models"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error.",
		})
		return
	}

	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
		})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event.",
		})
		return
	}

	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	var event models.Event

	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	userID := c.GetInt64("userID")
	event.UserID = userID

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error.",
		})
		return
	}

	c.JSON(http.StatusCreated, event)
}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
		})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the event.",
		})
		return
	}

	userID := c.GetInt64("userID")
	if event.UserID != userID {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "The event not found.",
		})
		return
	}

	var updatedEvent models.Event

	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}	

	updatedEvent.ID = eventId
	updatedEvent.UserID = userID

	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully.",
	})
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
		})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the event.",
		})
		return
	}

	userID := c.GetInt64("userID")
	if event.UserID != userID {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "The event not found.",
		})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully.",
	})
}
