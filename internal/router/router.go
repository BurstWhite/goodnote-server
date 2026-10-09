package router

import (
	"github.com/BurstWhite/goodnote-server/internal/note"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *note.Handler) *gin.Engine {
	r := gin.Default()

	notes := r.Group("/api/v1/notes")
	notes.GET("", handler.GetNotesHandler)
	notes.GET("/:id", handler.GetNoteByIDHandler)
	notes.POST("", handler.CreateNoteHandler)
	notes.PATCH("/:id", handler.ModifyNoteHandler)
	notes.DELETE("/:id", handler.DeleteNoteHandler)

	return r
}
