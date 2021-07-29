package note

import (
	"time"

	"github.com/RegisMelgaco/whathaveidone/whid/domain/shared"
)

type Note struct {
	Id          int
	Description string
	IsDeleted   bool
	CreatedAt   time.Time
	DeletedAt   time.Time
}

func NewNote(timeProvicer shared.TimeProvider, description string) (*Note, error) {
	if len(description) == 0 {
		return nil, ErrNoteMustHaveADescription
	}

	return &Note{
		Description: description,
		CreatedAt:   time.Now(),
	}, nil
}
