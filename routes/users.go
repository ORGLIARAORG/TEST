package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jarqvi/gin-crud/models"
	"github.com/jarqvi/gin-crud/utils"
)

func signUp(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully.",
	})
}

func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request.",
		})

		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials.",
		})

		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not authenticated user.",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully.",
		"token":   token,
	})
}
