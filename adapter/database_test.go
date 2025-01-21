package adapter_test

import (
	"os"
	"testing"

	"github.com/choonhong/hotel-data-merge/adapter"
	"github.com/stretchr/testify/require"
)

func TestConnet(t *testing.T) {
	_, err := adapter.Connect()
	defer os.Remove("gorm.db")

	require.NoError(t, err)
}
