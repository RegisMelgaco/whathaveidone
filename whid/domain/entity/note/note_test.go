package note_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/RegisMelgaco/whathaveidone/whid/domain/entity/note"
	"github.com/RegisMelgaco/whathaveidone/whid/domain/shared"
	"github.com/stretchr/testify/assert"
)

func TestNewNote(t *testing.T) {
	t.Parallel()

	now := time.Now()
	var timeProvider shared.TimeProvider = &shared.TimeProviderMock{
		NowFunc: func() time.Time {
			return now
		},
	}

	cases := []struct {
		testName    string
		description string
		expected    *note.Note
		expectedErr error
	}{
		{
			testName: "new note with description should return note",

			description: "Buy some bread.",

			expected: &note.Note{
				Id:          0,
				Description: "Buy some bread.",
				IsDeleted:   false,
				CreatedAt:   time.Now(),
				DeletedAt:   time.Time{},
			},
			expectedErr: nil,
		},
		{
			testName: fmt.Sprintf("new note without description should return error %v", note.ErrNoteMustHaveADescription),

			description: "",

			expected:    nil,
			expectedErr: note.ErrNoteMustHaveADescription,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			t.Parallel()

			actual, err := note.NewNote(timeProvider, c.description)

			assert.Equal(t, c.expected, actual)
			assert.Equal(t, c.expectedErr, err)
		})
	}
}
