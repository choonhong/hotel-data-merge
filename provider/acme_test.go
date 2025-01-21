package provider_test

import (
	"context"
	"testing"

	"github.com/choonhong/hotel-data-merge/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcme(t *testing.T) {
	subject := provider.Acme{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme"}

	hotels, err := subject.FetchAll(context.Background())
	require.NoError(t, err)

	require.Len(t, hotels, 3)

	assert.Equal(t, "iJhz", hotels[0].ID)
	assert.Equal(t, 5432, hotels[0].DestinationID)
	assert.Equal(t, "Beach Villas Singapore", hotels[0].Name)
	assert.Equal(t, 1.264751, *hotels[0].Latitude)
	assert.Equal(t, 103.824006, *hotels[0].Longitude)
	assert.Equal(t, "8 Sentosa Gateway, Beach Villas", hotels[0].Address)
	assert.Equal(t, "Singapore", hotels[0].City)
	assert.Equal(t, "SG", hotels[0].Country)
	assert.Equal(t, "098269", hotels[0].PostalCode)
	assert.Equal(t, "This 5 star hotel is located on the coastline of Singapore.", hotels[0].Description)
	assert.Equal(t, []string{"Pool", "BusinessCenter", "WiFi", "DryCleaning", "Breakfast"}, hotels[0].Amenities)

	assert.Equal(t, "SjyX", hotels[1].ID)
	assert.Equal(t, 5432, hotels[1].DestinationID)
	assert.Equal(t, "InterContinental Singapore Robertson Quay", hotels[1].Name)
	assert.Nil(t, hotels[1].Latitude)
	assert.Nil(t, hotels[1].Longitude)
	assert.Equal(t, "1 Nanson Road", hotels[1].Address)
	assert.Equal(t, "Singapore", hotels[1].City)
	assert.Equal(t, "SG", hotels[1].Country)
	assert.Equal(t, "238909", hotels[1].PostalCode)
	assert.Equal(t, "Enjoy sophisticated waterfront living at the new InterContinental® Singapore Robertson Quay, luxury's preferred address nestled in the heart of Robertson Quay along the Singapore River, with the CBD just five minutes drive away. Magnifying the comforts of home, each of our 225 studios and suites features a host of thoughtful amenities that combine modernity with elegance, whilst maintaining functional practicality. The hotel also features a chic, luxurious Club InterContinental Lounge.", hotels[1].Description)
	assert.Equal(t, []string{"Pool", "WiFi", "Aircon", "BusinessCenter", "BathTub", "Breakfast", "DryCleaning", "Bar"}, hotels[1].Amenities)

	assert.Equal(t, "f8c9", hotels[2].ID)
	assert.Equal(t, 1122, hotels[2].DestinationID)
	assert.Equal(t, "Hilton Shinjuku Tokyo", hotels[2].Name)
	assert.Nil(t, hotels[2].Latitude)
	assert.Nil(t, hotels[2].Longitude)
	assert.Equal(t, "160-0023, SHINJUKU-KU, 6-6-2 NISHI-SHINJUKU, JAPAN", hotels[2].Address)
	assert.Equal(t, "Tokyo", hotels[2].City)
	assert.Equal(t, "JP", hotels[2].Country)
	assert.Equal(t, "160-0023", hotels[2].PostalCode)
	assert.Equal(t, "Hilton Tokyo is located in Shinjuku, the heart of Tokyo's business, shopping and entertainment district, and is an ideal place to experience modern Japan. A complimentary shuttle operates between the hotel and Shinjuku station and the Tokyo Metro subway is connected to the hotel. Relax in one of the modern Japanese-style rooms and admire stunning city views. The hotel offers WiFi and internet access throughout all rooms and public space.", hotels[2].Description)
	assert.Equal(t, []string{"Pool", "WiFi", "BusinessCenter", "DryCleaning", "Breakfast", "Bar", "BathTub"}, hotels[2].Amenities)
}
