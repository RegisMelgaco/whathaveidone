package note

import "context"

//go:generate moq -out repository_mock.go . NoteRepo

type NoteRepo interface {
	CreateAccount(context.Context, *Note) error
}
