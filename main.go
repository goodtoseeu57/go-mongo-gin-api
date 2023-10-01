package main

import (
	"my-api/module/config"
	"my-api/module/handlers/auth"
	"my-api/module/handlers/noteshandler"
	authMiddleware "my-api/module/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadDBConfig()
	config.FirebaseInit()
	r := gin.Default()

	r.POST("/create-note", authMiddleware.FirebaseAuthMiddleware(), noteshandler.CreateNote)
	r.GET("/notes/:id", authMiddleware.FirebaseAuthMiddleware(), noteshandler.GetNoteById)
	r.GET("notes", authMiddleware.FirebaseAuthMiddleware(), noteshandler.GetNotes)

	r.POST("register", auth.Register)

	r.Run(":8000")
}
