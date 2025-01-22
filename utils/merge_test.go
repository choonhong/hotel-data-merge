package utils_test

import (
	"testing"

	"github.com/choonhong/hotel-data-merge/utils"
	"github.com/stretchr/testify/assert"
)

func TestFindMostAverageString(t *testing.T) {
	t.Run("no string", func(t *testing.T) {
		countries := []string{}
		assert.Equal(t, "", utils.FindMostAverageString(countries))
	})

	t.Run("1 string", func(t *testing.T) {
		countries := []string{"Singapore"}
		assert.Equal(t, "Singapore", utils.FindMostAverageString(countries))
	})

	t.Run("2 strings", func(t *testing.T) {
		countries := []string{"SG", "Singapore"}
		assert.Equal(t, "Singapore", utils.FindMostAverageString(countries))
	})

	t.Run("3 strings", func(t *testing.T) {
		countries := []string{"SG", "Singapore", "singapore"}
		assert.Equal(t, "Singapore", utils.FindMostAverageString(countries))
	})

	t.Run("3 strings different order", func(t *testing.T) {
		countries := []string{"singapore", "SG", "Singapore"}
		assert.Equal(t, "Singapore", utils.FindMostAverageString(countries))
	})
}
