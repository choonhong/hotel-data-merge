package internal_test

import (
	"context"
	"os"
	"testing"

	"github.com/choonhong/hotel-data-merge/adapter"
	"github.com/choonhong/hotel-data-merge/internal"
	"github.com/choonhong/hotel-data-merge/provider"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func TestFetchAndMergeHotels(t *testing.T) {
	db, err := adapter.Connect()
	defer os.Remove("gorm.db")

	require.NoError(t, err)

	subject := internal.HotelService{
		HotelRepo: &adapter.HotelRepository{DB: db},
		Providers: []internal.Provider{
			&provider.Acme{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme"},
			&provider.Paperflies{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies"},
			&provider.Patagonia{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia"},
		},
	}

	hotels, err := subject.FetchAndMergeHotels(context.Background())
	require.NoError(t, err)

	spew.Dump(hotels)
}
