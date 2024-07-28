package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jarqvi/gin-crud/models"
)

func registerForEvent(c *gin.Context) {
	userID := c.GetInt64("userID")

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

	err = event.Register(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not register user for event.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registered!",
	})
}

func cancelRegistration(c *gin.Context) {
	userID := c.GetInt64("userID")

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
		})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not cancel registration.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cancel registration successfully!",
	})
}
