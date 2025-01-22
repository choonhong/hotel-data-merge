package internal_test

import (
	"context"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/choonhong/hotel-data-merge/adapter"
	"github.com/choonhong/hotel-data-merge/database"
	"github.com/choonhong/hotel-data-merge/ent"
	"github.com/choonhong/hotel-data-merge/internal"
	"github.com/choonhong/hotel-data-merge/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchAndMergeHotels(t *testing.T) {
	// Connect to database
	db, err := database.Connect()
	require.NoError(t, err)
	defer db.Close()

	// Create hotel service
	subject := internal.HotelService{
		HotelRepo: &adapter.HotelRepository{Client: db},
		Providers: []internal.Provider{
			&provider.Acme{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme"},
			&provider.Paperflies{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies"},
			&provider.Patagonia{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia"},
		},
	}

	// Fetch and merge hotels
	hotels, err := subject.FetchAndMergeHotels(context.Background())
	require.NoError(t, err)
	require.Len(t, hotels, 3)

	// Check the result of merging
	hotel := hotels["iJhz"]
	assert.Equal(t, "iJhz", hotel.ID)
	assert.Equal(t, 5432, hotel.DestinationID)
	assert.Equal(t, "Beach Villas Singapore", hotel.Name)
	assert.Equal(t, 1.264751, *hotel.Latitude)
	assert.Equal(t, 103.824006, *hotel.Longitude)
	assert.Equal(t, "8 Sentosa Gateway, Beach Villas, 098269", hotel.Address)
	assert.Equal(t, "Singapore", hotel.City)
	assert.Equal(t, "Singapore", hotel.Country)
	assert.Equal(t, "098269", hotel.PostalCode)
	assert.Len(t, hotel.Amenities, 14)
	assert.True(t, reflect.DeepEqual([]string{
		"outdoor pool",
		"business center",
		"wifi",
		"drycleaning",
		"breakfast",
		"indoor pool",
		"childcare",
		"tv",
		"coffee machine",
		"kettle",
		"hair dryer",
		"iron",
		"aircon",
		"tub",
	}, hotel.Amenities))
	assert.Len(t, hotel.Images, 6)
	assert.Len(t, hotel.BookingConditions, 5)
}

func TestMergeHotels(t *testing.T) {
	hotels := []*ent.Hotel{}

	// Load test data from acme.json
	data, err := os.ReadFile("../testdata/acme.json")
	require.NoError(t, err)
	acmeHotels := []provider.AcmeData{}
	require.NoError(t, json.Unmarshal(data, &acmeHotels))
	for _, acmeHotel := range acmeHotels {
		hotels = append(hotels, acmeHotel.ToHotel())
	}

	// Load test data from patagonia.json
	data, err = os.ReadFile("../testdata/paperflies.json")
	require.NoError(t, err)
	paperfliesHotels := []provider.PaperfliesData{}
	require.NoError(t, json.Unmarshal(data, &paperfliesHotels))
	for _, paperfliesHotel := range paperfliesHotels {
		hotels = append(hotels, paperfliesHotel.ToHotel())
	}

	// Merge hotels
	mergedHotels := internal.MergeHotels(hotels)
	require.Len(t, mergedHotels, 3)

	// Check the result of merging
	hotel := mergedHotels["iJhz"]
	assert.Equal(t, "iJhz", hotel.ID)
	assert.Equal(t, 5432, hotel.DestinationID)
	assert.Equal(t, "Beach Villas Singapore", hotel.Name)
	assert.Equal(t, 1.264751, *hotel.Latitude)
	assert.Equal(t, 103.824006, *hotel.Longitude)
	assert.Equal(t, "8 Sentosa Gateway, Beach Villas, 098269", hotel.Address)
	assert.Equal(t, "Singapore", hotel.City)
	assert.Equal(t, "Singapore", hotel.Country)
	assert.Equal(t, "098269", hotel.PostalCode)
	assert.Len(t, hotel.Amenities, 12)
	assert.True(t, reflect.DeepEqual([]string{
		"outdoor pool",
		"business center",
		"wifi",
		"drycleaning",
		"breakfast",
		"indoor pool",
		"childcare",
		"tv",
		"coffee machine",
		"kettle",
		"hair dryer",
		"iron",
	}, hotel.Amenities))
	assert.Len(t, hotel.Images, 3)
	assert.Len(t, hotel.BookingConditions, 5)
}
