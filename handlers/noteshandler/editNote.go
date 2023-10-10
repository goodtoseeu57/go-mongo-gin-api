package noteshandler

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"my-api/module/config"
	"my-api/module/models"
	"net/http"
)

func EditNote(c *gin.Context) {
	userId, exists := c.Get("token")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Permission denied"})
	}

	var note models.Note

	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	filter := bson.M{"_id": note.ID, "user_id": userId}
	update := bson.M{"$set": note}

	result, err := config.NotesCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating note"})
		return
	}
	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note updated successfully"})
}
