package expense

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	expense := Expense{
		ID:          1,
		Amount:      10,
		Category:    "Food",
		Description: "Lunch",
		Date:        time.Date(2025, time.April, 15, 0, 0, 0, 0, time.UTC),
	}
	record := encode(expense)
	want := []string{
		"1",
		"10",
		"Food",
		"Lunch",
		"2025-04-15",
	}
	assert.Equal(t, want, record)
}

func TestDecode(t *testing.T) {
	t.Run("invalid id", func(t *testing.T) {
		expense, err := decode([]string{"invalid", "10", "Food", "Lunch", "2025-04-15"})
		assert.Error(t, err)
		assert.Nil(t, expense)
	})
	t.Run("invalid amount", func(t *testing.T) {
		expense, err := decode([]string{"1", "invalid", "Food", "Lunch", "2025-04-15"})
		assert.Error(t, err)
		assert.Nil(t, expense)
	})
	t.Run("invalid date", func(t *testing.T) {
		expense, err := decode([]string{"1", "10", "Food", "Lunch", "invalid"})
		assert.Error(t, err)
		assert.Nil(t, expense)
	})
	t.Run("slice length not 5", func(t *testing.T) {
		expense, err := decode([]string{"1", "10", "Food", "Lunch", "2025-10-10", ""})
		assert.Error(t, err)
		assert.Nil(t, expense)
	})

	t.Run("successful decoding", func(t *testing.T) {
		expense, err := decode([]string{
			"1",
			"10",
			"Food",
			"Lunch",
			"2025-04-15",
		})
		require.NoError(t, err)
		want := Expense{
			ID:          1,
			Amount:      10,
			Category:    "Food",
			Description: "Lunch",
			Date:        time.Date(2025, time.April, 15, 0, 0, 0, 0, time.UTC),
		}
		assert.Equal(t, want, *expense)
	})
}
