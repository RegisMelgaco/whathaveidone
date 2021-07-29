package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/RegisMelgaco/whathaveidone/whid/domain/entity/note"
	"github.com/RegisMelgaco/whathaveidone/whid/domain/shared"
	"github.com/RegisMelgaco/whathaveidone/whid/domain/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCreateNote(t *testing.T) {
	t.Parallel()

	now := time.Now()
	var timeProvider shared.TimeProvider = &shared.TimeProviderMock{
		NowFunc: func() time.Time {
			return now
		},
	}

	workingRepo := &note.NoteRepoMock{
		CreateAccountFunc: func(contextMoqParam context.Context, note *note.Note) error {
			return nil
		},
	}

	dbErr := errors.New("Some sort of db problem")

	cases := []struct {
		testName string

		repo        *note.NoteRepoMock
		description string

		expected             *note.Note
		expectedErr          error
		isExpectedToCallRepo bool
	}{
		{
			testName: "Create note with description and working repo should call repo and retrieve note",

			repo:        workingRepo,
			description: "Walk to the sunset",

			expected: &note.Note{
				Description: "Walk to the sunset",
				CreatedAt:   timeProvider.Now(),
			},
			expectedErr:          nil,
			isExpectedToCallRepo: true,
		},
		{
			testName: "Create note with empty description should not call repo and retrieve error",

			repo:        workingRepo,
			description: "",

			expected:             nil,
			expectedErr:          note.ErrNoteMustHaveADescription,
			isExpectedToCallRepo: false,
		},
		{
			testName: "Create note with description and not working repo should call repo and retrive error",

			repo: &note.NoteRepoMock{
				CreateAccountFunc: func(contextMoqParam context.Context, note *note.Note) error {
					return dbErr
				},
			},
			description: "Wash some dishes",

			expected:             nil,
			expectedErr:          dbErr,
			isExpectedToCallRepo: true,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			t.Parallel()

			var uc usecase.NoteUsecase = usecase.NoteUsecaseImpl{
				Repo:         c.repo,
				TimeProvider: timeProvider,
			}

			actual, err := uc.CreateNote(context.Background(), c.description)

			assert.Equal(t, c.expected, actual)
			assert.Equal(t, c.expectedErr, err)

			wasRepoCalled := len(c.repo.CreateAccountCalls()) > 0
			assert.Equal(t, c.isExpectedToCallRepo, wasRepoCalled)
		})
	}
}
