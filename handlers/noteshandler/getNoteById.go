package noteshandler

import (
	"context"
	"fmt"
	"my-api/module/config"
	"my-api/module/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNoteById(c *gin.Context) {
	paramId := c.Param("id")

	userId, exists := c.Get("token")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Permission denied"})
		return
	}

	id, _ := primitive.ObjectIDFromHex(paramId)

	fmt.Printf("%v", id)

	filter := bson.M{"_id": id, "user_id": userId}

	var result models.Note

	err := config.NotesCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while finding a note"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"note": result,
	})
}
