package v1

import (
	"gin/forms"
	"gin/models"
	"log"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

//RegisterController ...
type RegisterController struct{}

//RegisterWithEmail ...
func (r RegisterController) RegisterWithEmail(c *gin.Context) {
	var f forms.StudentRegister
	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client, err := models.GetAuthClient(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO create user
	params := (&auth.UserToCreate{}).
		Email(f.Email).
		EmailVerified(false).
		// PhoneNumber("+15555550100").
		Password(f.Password).
		// DisplayName("John Doe").
		// PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)
	u, err := client.CreateUser(c, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", u)
	// TODO send verify email

	c.JSON(http.StatusOK, gin.H{"data": u})
}
