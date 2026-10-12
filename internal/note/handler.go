package note

import (
	"errors"
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

	note, err := h.service.GetNoteByID(c.Request.Context(), int64(id))
	switch {
	case err == nil:
		c.JSON(http.StatusOK, note)

	case errors.Is(err, ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{
			"error": "note not found",
		})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	}
}

func (h *Handler) CreateNoteHandler(c *gin.Context) {
	var req CreateNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	id, err := h.service.CreateNote(c.Request.Context(), req)
	switch {
	case err == nil:
		c.JSON(http.StatusCreated, gin.H{
			"id": id,
		})

	case errors.Is(err, ErrInvalidTitle) || errors.Is(err, ErrInvalidContent):
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	}

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
