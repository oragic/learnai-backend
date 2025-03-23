package http

import (
	"errors"
	"learnai/internal/core/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

type flashCardResponse struct {
	ID        uint64    `json:"id" example:"1"`
	UserID    uint64    `json:"userid" example:"1"`
	AudioFile string    `json:"audiofile" example:"Technology Flash Card"`
	Text      string    `json:"text" example:"Technology Flash Card"`
	CreatedAt time.Time `json:"created_at" example:"1970-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"1970-01-01T00:00:00Z"`
}

func newFlashCardResponse(flashCard *domain.Flashcard) flashCardResponse {
	return flashCardResponse{
		ID:        flashCard.ID,
		UserID:    flashCard.UserID,
		AudioFile: flashCard.AudioFile,
		Text:      flashCard.Text,
		CreatedAt: flashCard.CreatedAt,
		UpdatedAt: flashCard.UpdateAt,
	}
}

var errorStatusMap = map[error]int{
	domain.ErrInternal:                   http.StatusInternalServerError,
	domain.ErrDataNotFound:               http.StatusNotFound,
	domain.ErrConflictingData:            http.StatusConflict,
	domain.ErrInvalidCredentials:         http.StatusUnauthorized,
	domain.ErrUnauthorized:               http.StatusUnauthorized,
	domain.ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	domain.ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	domain.ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	domain.ErrInvalidToken:               http.StatusUnauthorized,
	domain.ErrExpiredToken:               http.StatusUnauthorized,
	domain.ErrForbidden:                  http.StatusForbidden,
	domain.ErrNoUpdatedData:              http.StatusBadRequest,
	domain.ErrInsufficientStock:          http.StatusBadRequest,
	domain.ErrInsufficientPayment:        http.StatusBadRequest,
}

func validationError(ctx *gin.Context, err error) {
	errMsgs := parseError(err)
	errRsp := newErrorResponse(errMsgs)
	ctx.JSON(http.StatusBadRequest, errRsp)
}

func handleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)
	ctx.JSON(statusCode, errRsp)
}

func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

type errorResponse struct {
	Success bool     `json:"success"  example:"false"`
	Message []string `json:"message" example:"Error message 1, Error message 2"`
}

func newErrorResponse(errMsgs []string) errorResponse {
	return errorResponse{
		Success: false,
		Message: errMsgs,
	}
}

func handleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, rsp)
}
