package auth

import (
	"context"
	"net/http"

	"my-api/module/config"
	"my-api/module/models"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := (&auth.UserToCreate{}).
		Email(user.Email).
		Password(user.Password)
	fbUser, err := config.AuthClient.CreateUser(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering user in Firebase"})
		return
	}

	user.UserID = fbUser.UID

	_, err = config.UsersCollection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting user into the database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "userId": user.UserID})
}
