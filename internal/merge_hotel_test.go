package internal_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/choonhong/hotel-data-merge/adapter"
	"github.com/choonhong/hotel-data-merge/database"
	"github.com/choonhong/hotel-data-merge/internal"
	"github.com/choonhong/hotel-data-merge/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchAndMergeHotels(t *testing.T) {
	db, err := database.Connect()
	require.NoError(t, err)
	defer db.Close()

	subject := internal.HotelService{
		HotelRepo: &adapter.HotelRepository{Client: db},
		Providers: []internal.Provider{
			&provider.Acme{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme"},
			&provider.Paperflies{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies"},
			&provider.Patagonia{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia"},
		},
	}

	hotels, err := subject.FetchAndMergeHotels(context.Background())
	require.NoError(t, err)

	for _, hotel := range hotels {
		fmt.Println("1. ID: " + hotel.ID)
		fmt.Printf("2. DestinationID: %d\n", hotel.DestinationID)
		fmt.Println("3. Name: " + hotel.Name)

		if hotel.Latitude != nil {
			fmt.Printf("4. Latitude: %f\n", *hotel.Latitude)
		}
		if hotel.Longitude != nil {
			fmt.Printf("5. Longitude: %f\n", *hotel.Longitude)
		}

		fmt.Println("6. Address: " + hotel.Address)
		fmt.Println("7. City: " + hotel.City)
		fmt.Println("8. Country: " + hotel.Country)
		fmt.Println("9. PostalCode: " + hotel.PostalCode)
		fmt.Println("10. Description: " + hotel.Description)
		fmt.Println("11. Amenities: " + strings.Join(hotel.Amenities, ", "))
		fmt.Printf("12. Images: %v\n", hotel.Images)
		fmt.Printf("13. Bookings: %v\n", hotel.BookingConditions)
		fmt.Println("=====================================")
	}
}

func TestFindMostAverageString(t *testing.T) {
	t.Run("no string", func(t *testing.T) {
		countries := []string{}
		assert.Equal(t, "", internal.FindMostAverageString(countries))
	})

	t.Run("1 string", func(t *testing.T) {
		countries := []string{"Singapore"}
		assert.Equal(t, "Singapore", internal.FindMostAverageString(countries))
	})

	t.Run("2 strings", func(t *testing.T) {
		countries := []string{"SG", "Singapore"}
		assert.Equal(t, "Singapore", internal.FindMostAverageString(countries))
	})

	t.Run("3 strings", func(t *testing.T) {
		countries := []string{"SG", "Singapore", "singapore"}
		assert.Equal(t, "Singapore", internal.FindMostAverageString(countries))
	})

	t.Run("3 strings different order", func(t *testing.T) {
		countries := []string{"singapore", "SG", "Singapore"}
		assert.Equal(t, "Singapore", internal.FindMostAverageString(countries))
	})
}