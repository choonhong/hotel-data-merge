package database_test

import (
	"testing"

	"github.com/choonhong/hotel-data-merge/database"
	"github.com/stretchr/testify/require"
)

func TestConnet(t *testing.T) {
	_, err := database.Connect()

	require.NoError(t, err)
}
