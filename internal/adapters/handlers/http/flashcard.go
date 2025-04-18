package http

import (
	"learnai/internal/core/domain"
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
	AudioFile string `json:"AudioFile" binding:"required" example:"Technology Flash Card"`
	Text      string `json:"text" binding:"required" example:"Technology Flash Card"`
}

func (h *FlashCardHandler) Create(ctx *gin.Context) {
	var request createRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, err)
		return
	}

	user := domain.Flashcard{
		AudioFile: request.AudioFile,
		Text:      request.Text,
	}

	_, err := h.svc.CreateFlashCard(ctx, &user)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newFlashCardResponse(&user)

	handleSuccess(ctx, rsp)

}

type getFlashCardRequest struct {
}

func (h *getFlashCardRequest) GetFlashCard(ctx *gin.Context) {

}
