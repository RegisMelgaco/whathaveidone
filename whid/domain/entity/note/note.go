package note

import "time"

type Note struct {
	Id          int
	Description string
	IsDeleted   bool
	CreatedAt   time.Time
	DeletedAt   time.Time
}

func NewNote(description string) (*Note, error) {
	return &Note{}, nil
}
