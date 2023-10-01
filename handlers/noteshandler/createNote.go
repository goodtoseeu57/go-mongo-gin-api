package noteshandler

import (
	"context"
	"my-api/module/config"
	"my-api/module/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {

	userId, exists := c.Get("token")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Permission denied"})
		return
	}

	var note models.Note

	note.UserId = userId.(string)

	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := config.NotesCollection.InsertOne(context.TODO(), note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting Note"})
		return
	}

	c.JSON(http.StatusCreated, result)

}
