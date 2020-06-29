package controllers

import (
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Test(c *gin.Context) {
	client, err := models.GetAuthClient(c)

	// fmt.Println(client, err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	user, err := client.GetUser(c, "d4aJk4FJOga9L8LNYNVnGLSAsYg1")
	if err != nil {

	}

	c.JSON(http.StatusBadRequest, gin.H{"message": user})
}
