package expense

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateID(t *testing.T) {
	t.Run("fails if id is negative", func(t *testing.T) {
		assert.Error(t, ValidateID(-1))
	})
	t.Run("fails if id is zero", func(t *testing.T) {
		assert.Error(t, ValidateID(0))
	})
	t.Run("validates if id is positive", func(t *testing.T) {
		assert.NoError(t, ValidateID(1))
	})
}

func TestValidateAmount(t *testing.T) {
	t.Run("fails if id is negative", func(t *testing.T) {
		assert.Error(t, ValidateAmount(-1))
	})
	t.Run("fails if id is zero", func(t *testing.T) {
		assert.Error(t, ValidateAmount(0))
	})
	t.Run("validates if id is positive", func(t *testing.T) {
		assert.NoError(t, ValidateAmount(1))
	})
}

func TestParseCategory(t *testing.T) {
	t.Run("fails with empty string", func(t *testing.T) {
		c, err := ParseCategory("")
		assert.Error(t, err)
		assert.Equal(t, "", c)
	})
	t.Run("fails with space only", func(t *testing.T) {
		c, err := ParseCategory("     ")
		assert.Error(t, err)
		assert.Equal(t, "", c)
	})
	t.Run("trims spaces on success", func(t *testing.T) {
		c, err := ParseCategory("  food  ")
		assert.NoError(t, err)
		assert.Equal(t, "food", c)
	})
}
func TestParseDescription(t *testing.T) {
	t.Run("fails with empty string", func(t *testing.T) {
		d, err := ParseDescription("")
		assert.Error(t, err)
		assert.Equal(t, "", d)
	})
	t.Run("fails with space only", func(t *testing.T) {
		d, err := ParseDescription("     ")
		assert.Error(t, err)
		assert.Equal(t, "", d)
	})
	t.Run("trims spaces on success", func(t *testing.T) {
		d, err := ParseDescription("  food  ")
		assert.NoError(t, err)
		assert.Equal(t, "food", d)
	})
}
