package provider_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/choonhong/hotel-data-merge/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaperflies(t *testing.T) {
	subject := provider.Paperflies{URL: "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies"}

	hotels, err := subject.FetchAll(context.Background())
	require.NoError(t, err)
	require.Len(t, hotels, 3)

	assert.Equal(t, "iJhz", hotels[0].ID)
	assert.Equal(t, 5432, hotels[0].DestinationID)
	assert.Equal(t, "Beach Villas Singapore", hotels[0].Name)
	assert.Equal(t, "8 Sentosa Gateway, Beach Villas, 098269", hotels[0].Address)
	assert.Equal(t, "Singapore", hotels[0].Country)
	assert.Equal(t, "Surrounded by tropical gardens, these upscale villas in elegant Colonial-style buildings are part of the Resorts World Sentosa complex and a 2-minute walk from the Waterfront train station. Featuring sundecks and pool, garden or sea views, the plush 1- to 3-bedroom villas offer free Wi-Fi and flat-screens, as well as free-standing baths, minibars, and tea and coffeemaking facilities. Upgraded villas add private pools, fridges and microwaves; some have wine cellars. A 4-bedroom unit offers a kitchen and a living room. There's 24-hour room and butler service. Amenities include posh restaurant, plus an outdoor pool, a hot tub, and free parking.", hotels[0].Description)
	assert.True(t, reflect.DeepEqual([]string{"outdoor pool", "indoor pool", "business center", "childcare", "tv", "coffee machine", "kettle", "hair dryer", "iron"}, hotels[0].Amenities))
	assert.Len(t, hotels[0].Images, 3)
	assert.Equal(t, []string{
		"All children are welcome. One child under 12 years stays free of charge when using existing beds. One child under 2 years stays free of charge in a child's cot/crib. One child under 4 years stays free of charge when using existing beds. One older child or adult is charged SGD 82.39 per person per night in an extra bed. The maximum number of children's cots/cribs in a room is 1. There is no capacity for extra beds in the room.",
		"Pets are not allowed.",
		"WiFi is available in all areas and is free of charge.",
		"Free private parking is possible on site (reservation is not needed).",
		"Guests are required to show a photo identification and credit card upon check-in. Please note that all Special Requests are subject to availability and additional charges may apply. Payment before arrival via bank transfer is required. The property will contact you after you book to provide instructions. Please note that the full amount of the reservation is due before arrival. Resorts World Sentosa will send a confirmation with detailed payment information. After full payment is taken, the property's details, including the address and where to collect keys, will be emailed to you. Bag checks will be conducted prior to entry to Adventure Cove Waterpark. === Upon check-in, guests will be provided with complimentary Sentosa Pass (monorail) to enjoy unlimited transportation between Sentosa Island and Harbour Front (VivoCity). === Prepayment for non refundable bookings will be charged by RWS Call Centre. === All guests can enjoy complimentary parking during their stay, limited to one exit from the hotel per day. === Room reservation charges will be charged upon check-in. Credit card provided upon reservation is for guarantee purpose. === For reservations made with inclusive breakfast, please note that breakfast is applicable only for number of adults paid in the room rate. Any children or additional adults are charged separately for breakfast and are to paid directly to the hotel.",
	}, hotels[0].BookingConditions)

	assert.Equal(t, "SjyX", hotels[1].ID)
	assert.Equal(t, 5432, hotels[1].DestinationID)
	assert.Equal(t, "InterContinental", hotels[1].Name)
	assert.Equal(t, "1 Nanson Rd, Singapore 238909", hotels[1].Address)
	assert.Equal(t, "Singapore", hotels[1].Country)
	assert.Equal(t, "InterContinental Singapore Robertson Quay is luxury's preferred address offering stylishly cosmopolitan riverside living for discerning travelers to Singapore. Prominently situated along the Singapore River, the 225-room inspiring luxury hotel is easily accessible to the Marina Bay Financial District, Central Business District, Orchard Road and Singapore Changi International Airport, all located a short drive away. The hotel features the latest in Club InterContinental design and service experience, and five dining options including Publico, an Italian landmark dining and entertainment destination by the waterfront.", hotels[1].Description)
	assert.True(t, reflect.DeepEqual([]string{"outdoor pool", "business center", "childcare", "parking", "bar", "dry cleaning", "wifi", "breakfast", "concierge", "aircon", "minibar", "tv", "bathtub", "hair dryer"}, hotels[1].Amenities))
	assert.Len(t, hotels[1].Images, 6)

	assert.Equal(t, "f8c9", hotels[2].ID)
	assert.Equal(t, 1122, hotels[2].DestinationID)
	assert.Equal(t, "Hilton Tokyo", hotels[2].Name)
	assert.Equal(t, "160-0023, SHINJUKU-KU, 6-6-2 NISHI-SHINJUKU, JAPAN", hotels[2].Address)
	assert.Equal(t, "Japan", hotels[2].Country)
	assert.Equal(t, "This sleek high-rise property is 10 minutes' walk from Shinjuku train station, 6 minutes' walk from the Tokyo Metropolitan Government Building and 3 km from Yoyogi Park. The polished rooms offer Wi-Fi and flat-screen TVs, plus minibars, sitting areas, and tea and coffeemaking facilities. Suites add living rooms, and access to a club lounge serving breakfast and cocktails. A free shuttle to Shinjuku station is offered. There's a chic Chinese restaurant, a sushi bar, and a grill restaurant with an open kitchen, as well as an English pub and a hip cocktail lounge. Other amenities include a gym, rooftop tennis courts, and a spa with an indoor pool.", hotels[2].Description)
	assert.Equal(t, []string{"indoor pool", "business center", "wifi", "tv", "aircon", "minibar", "bathtub", "hair dryer"}, hotels[2].Amenities)
	assert.Len(t, hotels[2].Images, 3)
	assert.Equal(t, []string{
		"All children are welcome. One child under 6 years stays free of charge when using existing beds. There is no capacity for extra beds in the room.",
		"Pets are not allowed.",
		"Wired internet is available in the hotel rooms and charges are applicable. WiFi is available in the hotel rooms and charges are applicable.",
		"Private parking is possible on site (reservation is not needed) and costs JPY 1500 per day.",
		"When booking more than 9 rooms, different policies and additional supplements may apply.",
		"The hotel's free shuttle is offered from Bus Stop #21 in front of Keio Department Store at Shinjuku Station. It is available every 20-minutes from 08:20-21:40. The hotel's free shuttle is offered from the hotel to Shinjuku Train Station. It is available every 20-minutes from 08:12-21:52. For more details, please contact the hotel directly. At the Executive Lounge a smart casual dress code is strongly recommended. Attires mentioned below are strongly discouraged and may not permitted: - Night attire (slippers, Yukata robe, etc.) - Gym clothes/sportswear (Tank tops, shorts, etc.) - Beachwear (flip-flops, sandals, etc.) and visible tattoos. Please note that due to renovation works, the Executive Lounge will be closed from 03 January 2019 until late April 2019. During this period, guests may experience some noise or minor disturbances. Smoking preference is subject to availability and cannot be guaranteed.",
	}, hotels[2].BookingConditions)
}
