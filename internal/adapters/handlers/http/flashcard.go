package http

import (
	"learnai/internal/core/port"

	"github.com/gin-gonic/gin"
)

type FlashCardHandler struct {
	svc port.FlashcardService
}

func NewFlashCardHandler(svc port.FlashcardService) *FlashCardHandler {
	return &FlashCardHandler{svc}
}

type createRequest struct {
}

func (h *FlashCardHandler) Create(ctx *gin.Context) {
}

type getFlashCardRequest struct {
}

func (h *getFlashCardRequest) GetFlashCard(ctx *gin.Context) {

}
