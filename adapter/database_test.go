package adapter_test

import (
	"testing"

	"github.com/choonhong/hotel-data-merge/adapter"
	"github.com/stretchr/testify/require"
)

func TestConnet(t *testing.T) {
	_, err := adapter.Connect()

	require.NoError(t, err)
}
