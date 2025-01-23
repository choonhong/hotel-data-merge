package utils_test

import (
	"testing"

	"github.com/choonhong/hotel-data-merge/utils"
	"github.com/stretchr/testify/assert"
)

func TestCalculateAverage(t *testing.T) {
	t.Run("no number", func(t *testing.T) {
		numbers := []float64{}
		assert.Nil(t, utils.CalculateAverage(numbers))
	})

	t.Run("1 number", func(t *testing.T) {
		numbers := []float64{1.0}
		assert.Equal(t, 1.0, *utils.CalculateAverage(numbers))
	})

	t.Run("2 numbers", func(t *testing.T) {
		numbers := []float64{1.0, 2.0}
		assert.Equal(t, 1.5, *utils.CalculateAverage(numbers))
	})

	t.Run("3 numbers", func(t *testing.T) {
		numbers := []float64{1.0, 2.0, 3.0}
		assert.Equal(t, 2.0, *utils.CalculateAverage(numbers))
	})

	t.Run("3 numbers with negative", func(t *testing.T) {
		numbers := []float64{-1.0, 0.0, 1.0}
		assert.Equal(t, 0.0, *utils.CalculateAverage(numbers))
	})
}

func TestFindMostAverageString(t *testing.T) {
	t.Run("no string", func(t *testing.T) {
		strs := []string{}
		assert.Equal(t, "", utils.FindMostRepresentativeString(strs))
	})

	t.Run("1 string", func(t *testing.T) {
		strs := []string{"Singapore"}
		assert.Equal(t, "Singapore", utils.FindMostRepresentativeString(strs))
	})

	t.Run("2 strings", func(t *testing.T) {
		strs := []string{"SG", "Singapore"}
		assert.Equal(t, "Singapore", utils.FindMostRepresentativeString(strs))
	})

	t.Run("3 strings", func(t *testing.T) {
		strs := []string{"SG", "Singapore", "singapore"}
		assert.Equal(t, "Singapore", utils.FindMostRepresentativeString(strs))
	})

	t.Run("3 strings different order", func(t *testing.T) {
		strs := []string{"singapore", "SG", "Singapore"}
		assert.Equal(t, "Singapore", utils.FindMostRepresentativeString(strs))
	})
}

func TestGetUniqueStrings(t *testing.T) {
	t.Run("no string", func(t *testing.T) {
		strs := []string{}
		assert.Equal(t, []string{}, utils.DedupStrings(strs))
	})

	t.Run("1 string", func(t *testing.T) {
		strs := []string{"pool"}
		assert.Equal(t, []string{"pool"}, utils.DedupStrings(strs))
	})

	t.Run("2 strings", func(t *testing.T) {
		strs := []string{"pool", "wifi"}
		assert.Equal(t, []string{"pool", "wifi"}, utils.DedupStrings(strs))
	})

	t.Run("2 similar strings", func(t *testing.T) {
		strs := []string{"pool", "pook"}
		assert.Equal(t, []string{"pool"}, utils.DedupStrings(strs))
	})

	t.Run("iron and aircon", func(t *testing.T) {
		strs := []string{"iron", "aircon"}
		assert.Equal(t, []string{"iron", "aircon"}, utils.DedupStrings(strs))
	})

	t.Run("strings with substring", func(t *testing.T) {
		strs := []string{"pool", "wifi", "indoor pool"}
		assert.Equal(t, []string{"indoor pool", "wifi"}, utils.DedupStrings(strs))
	})
}
