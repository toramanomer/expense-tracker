package expense

import (
	"bytes"
	"errors"
	"io"
	"strings"
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
			Category:    "Food",
			Description: "Test Expense",
			Date:        time.Date(2025, time.April, 20, 0, 0, 0, 0, time.Local),
		}, w)
		require.NoError(t, err)
		assert.Equal(t, "1,10,Food,Test Expense,2025-04-20\n", w.String())

		err = s.add(Expense{
			ID:          2,
			Amount:      20,
			Category:    "Food",
			Description: "Test Expense 2",
			Date:        time.Date(2025, time.April, 21, 0, 0, 0, 0, time.Local),
		}, w)
		require.NoError(t, err)
		assert.Equal(t, "1,10,Food,Test Expense,2025-04-20\n2,20,Food,Test Expense 2,2025-04-21\n", w.String())
	})
}

func TestStorageFS_List(t *testing.T) {
	s := NewStorageFS(t.TempDir())
	exp1 := Expense{
		ID:          1,
		Amount:      10,
		Description: "Test Expense",
		Date:        time.Date(2025, time.April, 20, 0, 0, 0, 0, time.UTC),
	}
	exp2 := Expense{
		ID:          2,
		Amount:      20,
		Description: "Test Expense 2",
		Date:        time.Date(2025, time.April, 21, 0, 0, 0, 0, time.UTC),
	}

	w := bytes.NewBufferString("")
	s.add(exp1, w)
	s.add(exp2, w)

	expenses, err := s.list(w)
	require.NoError(t, err)
	assert.Equal(t, []Expense{exp1, exp2}, expenses)
}

type readerWriteTruncater struct {
	s           string
	i           int64 // current index
	readErr     error
	seekErr     error
	truncateErr error
	writeErr    error
}

func (rwt *readerWriteTruncater) Read(p []byte) (int, error) {
	if rwt.readErr != nil {
		return 0, rwt.readErr
	}

	if rwt.i >= int64(len(rwt.s)) {
		return 0, io.EOF
	}

	n := copy(p, rwt.s[rwt.i:])
	rwt.i += int64(n)

	return n, nil
}

func (rwt *readerWriteTruncater) Seek(offset int64, whence int) (int64, error) {
	if rwt.seekErr != nil {
		return 0, rwt.seekErr
	}

	var abs int64
	switch whence {
	case io.SeekStart:
		abs = offset
	case io.SeekCurrent:
		abs = rwt.i + offset
	case io.SeekEnd:
		abs = int64(len(rwt.s)) + offset
	default:
		return 0, errors.New("invalid whence")
	}

	if abs < 0 {
		return 0, errors.New("negative position")
	}

	rwt.i = abs
	return abs, nil
}

func (rwt *readerWriteTruncater) Write(p []byte) (int, error) {
	if rwt.writeErr != nil {
		return 0, rwt.writeErr
	}

	rwt.s = rwt.s[:rwt.i] + string(p) + rwt.s[rwt.i:]
	return len(p), nil
}

func (rwt *readerWriteTruncater) Truncate(size int64) error {
	if rwt.truncateErr != nil {
		return rwt.truncateErr
	}

	if size < 0 {
		return errors.New("negative size")
	}

	if size > int64(len(rwt.s)) {
		return errors.New("size exceeds current length")
	}

	rwt.s = rwt.s[:size]

	return nil
}

func TestStorageFS_Delete(t *testing.T) {
	exp1 := Expense{
		ID:          1,
		Amount:      10,
		Category:    "Food",
		Date:        time.Date(2025, time.April, 25, 0, 0, 0, 0, time.UTC),
		Description: "Lunch",
	}
	exp2 := Expense{
		ID:          2,
		Amount:      20,
		Category:    "Food",
		Date:        time.Date(2025, time.April, 25, 0, 0, 0, 0, time.UTC),
		Description: "Lunch",
	}
	t.Run("successful delete", func(t *testing.T) {
		rwt := &readerWriteTruncater{
			s: strings.Join([]string{
				strings.Join(encode(exp1), ","),
				strings.Join(encode(exp2), ",") + "\n",
			}, "\n"),
		}

		expected := strings.Join(encode(exp2), ",") + "\n"

		s := &StorageFS{}
		err := s.delete(1, rwt)
		assert.NoError(t, err)
		assert.Equal(t, expected, rwt.s)
	})

	t.Run("truncate err", func(t *testing.T) {
		str := strings.Join([]string{
			strings.Join(encode(exp1), ","),
			strings.Join(encode(exp2), ",") + "\n",
		}, "\n")
		rwt := &readerWriteTruncater{s: str}
		rwt.truncateErr = errors.New("write error")

		s := &StorageFS{}
		err := s.delete(1, rwt)
		assert.ErrorIs(t, err, rwt.truncateErr)
		assert.Equal(t, str, rwt.s)
	})

	t.Run("seek err", func(t *testing.T) {
		str := strings.Join([]string{
			strings.Join(encode(exp1), ","),
			strings.Join(encode(exp2), ",") + "\n",
		}, "\n")
		rwt := &readerWriteTruncater{s: str}
		rwt.seekErr = errors.New("write error")

		s := &StorageFS{}
		err := s.delete(1, rwt)
		assert.ErrorIs(t, err, rwt.seekErr)
		assert.Equal(t, str, rwt.s)
	})

	t.Run("not found", func(t *testing.T) {
		str := strings.Join([]string{
			strings.Join(encode(exp1), ","),
			strings.Join(encode(exp2), ",") + "\n",
		}, "\n")
		rwt := &readerWriteTruncater{s: str}

		s := &StorageFS{}
		err := s.delete(3, rwt)
		assert.ErrorIs(t, err, ErrExpenseNotFound)
		assert.Equal(t, str, rwt.s)
	})
}
