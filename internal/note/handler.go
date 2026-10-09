package note

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetNotesHandler(c *gin.Context) {
	notes, err := h.service.GetNotes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unknown server error happened",
		})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func (h *Handler) GetNoteByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id: must be an integer",
		})
		return
	}

	note, err := h.service.GetNotesByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Note of ID %v not found.", id),
		})
		return
	}

	c.JSON(http.StatusOK, note)
}

func (h *Handler) CreateNoteHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "API Not Implemented",
	})
}

func (h *Handler) ModifyNoteHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "API Not Implemented",
	})
}

func (h *Handler) DeleteNoteHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "API Not Implemented",
	})
}
