package usecase

import (
	"context"

	"github.com/RegisMelgaco/whathaveidone/whid/domain/entity/note"
	"github.com/RegisMelgaco/whathaveidone/whid/domain/shared"
)

type NoteUsecase interface {
	CreateNote(ctx context.Context, description string) (*note.Note, error)
}

type NoteUsecaseImpl struct {
	Repo         note.NoteRepo
	TimeProvider shared.TimeProvider
}

func (uc NoteUsecaseImpl) CreateNote(ctx context.Context, description string) (*note.Note, error) {
	note, err := note.NewNote(uc.TimeProvider, description)
	if err != nil {
		return nil, err
	}

	err = uc.Repo.CreateAccount(ctx, note)
	if err != nil {
		return nil, err
	}

	return note, nil
}
