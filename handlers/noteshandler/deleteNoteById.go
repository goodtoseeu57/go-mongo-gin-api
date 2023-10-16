package noteshandler

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"my-api/module/config"
	"net/http"
)

func deleteNoteById(c *gin.Context) {
	paramId := c.Param("id")
	userId, exists := c.Get("token")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Permission denied"})
	}

	id, _ := primitive.ObjectIDFromHex(paramId)

	filter := bson.M{"_id": id, "user_id": userId}

	result, err := config.NotesCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting Note"})
		return
	}

	c.JSON(http.StatusOK, result)
}
