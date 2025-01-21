package provider_test

import (
	"context"
	"testing"

	"github.com/choonhong/hotel-data-merge/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPatagonia(t *testing.T) {
	subject := provider.Patagonia{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia"}

	hotels, err := subject.FetchAll(context.Background())
	require.NoError(t, err)
	require.Len(t, hotels, 2)

	assert.Equal(t, "iJhz", hotels[0].ID)
	assert.Equal(t, 5432, hotels[0].DestinationID)
	assert.Equal(t, "Beach Villas Singapore", hotels[0].Name)
	assert.Equal(t, 1.264751, *hotels[0].Latitude)
	assert.Equal(t, 103.824006, *hotels[0].Longitude)
	assert.Equal(t, "8 Sentosa Gateway, Beach Villas, 098269", hotels[0].Address)
	assert.Equal(t, "Located at the western tip of Resorts World Sentosa, guests at the Beach Villas are guaranteed privacy while they enjoy spectacular views of glittering waters. Guests will find themselves in paradise with this series of exquisite tropical sanctuaries, making it the perfect setting for an idyllic retreat. Within each villa, guests will discover living areas and bedrooms that open out to mini gardens, private timber sundecks and verandahs elegantly framing either lush greenery or an expanse of sea. Guests are assured of a superior slumber with goose feather pillows and luxe mattresses paired with 400 thread count Egyptian cotton bed linen, tastefully paired with a full complement of luxurious in-room amenities and bathrooms boasting rain showers and free-standing tubs coupled with an exclusive array of ESPA amenities and toiletries. Guests also get to enjoy complimentary day access to the facilities at Asia’s flagship spa – the world-renowned ESPA.", hotels[0].Description)
	assert.Equal(t, []string{"aircon", "tv", "coffee machine", "kettle", "hair dryer", "iron", "tub"}, hotels[0].Amenities)
	assert.Len(t, hotels[0].Images, 4)
	assert.Equal(t, "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg", hotels[0].Images[0].URL)
	assert.Equal(t, "Double room", hotels[0].Images[0].Description)
	assert.Equal(t, "rooms", hotels[0].Images[0].Category)
	assert.Equal(t, "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/4.jpg", hotels[0].Images[1].URL)
	assert.Equal(t, "Bathroom", hotels[0].Images[1].Description)
	assert.Equal(t, "rooms", hotels[0].Images[0].Category)
	assert.Equal(t, "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/0.jpg", hotels[0].Images[2].URL)
	assert.Equal(t, "RWS", hotels[0].Images[2].Description)
	assert.Equal(t, "amenities", hotels[0].Images[2].Category)
	assert.Equal(t, "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/6.jpg", hotels[0].Images[3].URL)
	assert.Equal(t, "Sentosa Gateway", hotels[0].Images[3].Description)
	assert.Equal(t, "amenities", hotels[0].Images[3].Category)

	assert.Equal(t, "f8c9", hotels[1].ID)
	assert.Equal(t, 1122, hotels[1].DestinationID)
	assert.Equal(t, "Hilton Tokyo Shinjuku", hotels[1].Name)
	assert.Equal(t, 35.6926, *hotels[1].Latitude)
	assert.Equal(t, 139.690965, *hotels[1].Longitude)
	assert.Nil(t, hotels[1].Amenities)
	assert.Len(t, hotels[1].Images, 3)
	assert.Equal(t, "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i10_m.jpg", hotels[1].Images[0].URL)
	assert.Equal(t, "Suite", hotels[1].Images[0].Description)
	assert.Equal(t, "rooms", hotels[1].Images[0].Category)
	assert.Equal(t, "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i11_m.jpg", hotels[1].Images[1].URL)
	assert.Equal(t, "Suite - Living room", hotels[1].Images[1].Description)
	assert.Equal(t, "rooms", hotels[1].Images[1].Category)
	assert.Equal(t, "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i57_m.jpg", hotels[1].Images[2].URL)
	assert.Equal(t, "Bar", hotels[1].Images[2].Description)
	assert.Equal(t, "amenities", hotels[1].Images[2].Category)
}
