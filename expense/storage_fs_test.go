package expense

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorageFS_GenerateID(t *testing.T) {
	t.Run("first id is 1", func(t *testing.T) {
		s := &StorageFS{
			idsfile: filepath.Join(t.TempDir(), "ids.txt"),
		}
		id, err := s.GenerateID()
		assert.NoError(t, err)
		assert.Equal(t, 1, id)
	})

	t.Run("sequential ids", func(t *testing.T) {
		s := &StorageFS{
			idsfile: filepath.Join(t.TempDir(), "ids.txt"),
		}
		id1, err := s.GenerateID()
		assert.NoError(t, err)
		assert.Equal(t, 1, id1)

		id2, err := s.GenerateID()
		assert.NoError(t, err)
		assert.Equal(t, 2, id2)
	})
}
