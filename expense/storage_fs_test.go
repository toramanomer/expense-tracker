package expense

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorageFS_GenerateID(t *testing.T) {
	t.Run("first id is 1", func(t *testing.T) {
		s := NewStorageFS(t.TempDir())
		id, err := s.GenerateID()
		assert.NoError(t, err)
		assert.Equal(t, 1, id)
	})

	t.Run("sequential ids", func(t *testing.T) {
		s := NewStorageFS(t.TempDir())
		id1, err := s.GenerateID()
		assert.NoError(t, err)
		assert.Equal(t, 1, id1)

		id2, err := s.GenerateID()
		assert.NoError(t, err)
		assert.Equal(t, 2, id2)
	})
}

func TestStorageFS_Add(t *testing.T) {
	t.Run("adds a new record", func(t *testing.T) {
		s := NewStorageFS(t.TempDir())

		w := bytes.NewBufferString("")

		err := s.add(Expense{
			ID:          1,
			Amount:      10,
			Description: "Test Expense",
			Date:        time.Date(2025, time.April, 20, 0, 0, 0, 0, time.Local),
		}, w)
		require.NoError(t, err)
		assert.Equal(t, "1,Test Expense,10,2025-04-20\n", w.String())

		err = s.add(Expense{
			ID:          2,
			Amount:      20,
			Description: "Test Expense 2",
			Date:        time.Date(2025, time.April, 21, 0, 0, 0, 0, time.Local),
		}, w)
		require.NoError(t, err)
		assert.Equal(t, "1,Test Expense,10,2025-04-20\n2,Test Expense 2,20,2025-04-21\n", w.String())
	})
}
