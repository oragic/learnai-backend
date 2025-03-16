package domain

import (
	"time"
)

type Flashcard struct {
	ID        uint64
	UserID    uint64
	AudioFile string
	Text      string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type FlashcardSet struct {
	ID          uint64
	UserID      uint64
	Title       string
	Description string
	Flashcards  []Flashcard
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
