package port

import (
	"context"
	"learnai/internal/core/domain"
)

type FlashCardRepository interface {
	CreateFlashCard(ctx context.Context, flashcard *domain.Flashcard) (*domain.Flashcard, error)
	GetFlashCardByID(ctx context.Context, id uint64) (*domain.Flashcard, error)
	GetFlashCard(ctx context.Context, id uint64) ([]domain.Flashcard, error)
	DeleteFlashCard(ctx context.Context, id uint64) error
}

type Flashcard interface {
	CreateFlashCard(ctx context.Context, flashcard *domain.Flashcard) (*domain.Flashcard, error)
	GetFlashCard(ctx context.Context, id uint64) ([]domain.Flashcard, error)
	GetFlashCardByID(ctx context.Context, id uint64) (*domain.Flashcard, error)
	DeleteFlashCard(ctx context.Context, id uint64) error
}
