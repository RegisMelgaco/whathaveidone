package note_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/RegisMelgaco/whathaveidone/whid/domain/entity/note"
	"github.com/stretchr/testify/assert"
)

func TestNewNote(t *testing.T) {
	t.Parallel()

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

			actual, err := note.NewNote(c.description)

			if actual != nil && c.expected != nil {
				createdAtDelta := actual.CreatedAt.Sub(c.expected.CreatedAt)
				assert.LessOrEqual(t, createdAtDelta, time.Second)
				c.expected.CreatedAt = actual.CreatedAt
			}

			assert.Equal(t, c.expected, actual)
			assert.Equal(t, c.expectedErr, err)
		})
	}
}
