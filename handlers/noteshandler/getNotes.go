package noteshandler

import (
	"context"
	"fmt"
	"my-api/module/config"
	"my-api/module/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetNotes(c *gin.Context) {
	userId, exists := c.Get("token")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Permission denied"})
		return
	}

	fmt.Printf("%v", userId)

	filter := bson.M{"user_id": userId}

	cursor, err := config.NotesCollection.Find(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't retrieve notes"})
	}

	var results []models.Note
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"notes": results,
		},
	)
}
