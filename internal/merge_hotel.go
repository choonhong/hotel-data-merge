package internal

import (
	"context"
	"fmt"

	"github.com/choonhong/hotel-data-merge/ent"
	"github.com/choonhong/hotel-data-merge/ent/schema"
	"github.com/choonhong/hotel-data-merge/utils"
)

type HotelService struct {
	HotelRepo HotelRepo
	Providers []Provider
}

// FetchAndMergeHotels fetches hotels from all providers and merges them into a single hotel list.
func (s *HotelService) FetchAndMergeHotels(ctx context.Context) ([]*ent.Hotel, error) {
	// Fetch hotels from all providers
	hotelMap, err := s.fetchHotels(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetchHotels: %w", err)
	}

	// Merge hotels with the same ID
	hotels, err := s.mergeHotels(hotelMap)
	if err != nil {
		return nil, fmt.Errorf("mergeHotels: %w", err)
	}

	fmt.Println("Saving merged hotels...")

	// Save merged hotels
	for _, hotel := range hotels {
		fmt.Println("Merged hotel " + hotel.ID)

		if err := s.HotelRepo.Save(context.Background(), hotel); err != nil {
			return nil, fmt.Errorf("Save: %w", err)
		}
	}

	fmt.Println("Fetch and merge hotels completed.")

	return hotels, nil
}

// fetchHotels fetches hotels from all providers and groups them by ID.
func (s *HotelService) fetchHotels(ctx context.Context) (map[string][]*ent.Hotel, error) {
	hotelMap := map[string][]*ent.Hotel{}

	// Fetch hotels from all providers
	for _, provider := range s.Providers {
		fmt.Println("Fetching hotels from ", provider.Name())

		hotels, err := provider.FetchAll(ctx)
		if err != nil {
			return nil, fmt.Errorf("FetchAll: %w", err)
		}

		// Group hotels by ID
		for _, hotel := range hotels {
			if _, ok := hotelMap[hotel.ID]; !ok {
				hotelMap[hotel.ID] = []*ent.Hotel{}
			}

			hotelMap[hotel.ID] = append(hotelMap[hotel.ID], hotel)
		}
	}

	return hotelMap, nil
}

// mergeHotels merges hotels with the same ID into a single hotel list.
func (s *HotelService) mergeHotels(hotelMap map[string][]*ent.Hotel) ([]*ent.Hotel, error) {
	fmt.Println("Merging hotels...")

	unifiedHotels := []*ent.Hotel{}

	// Merge hotels with the same ID
	for _, hotels := range hotelMap {
		if len(hotels) == 1 {
			unifiedHotels = append(unifiedHotels, hotels[0])
			continue
		}

		// Prepare data for merging
		names := []string{}
		latitudes := []float64{}
		longitudes := []float64{}
		longestAddress := ""
		cities := []string{}
		countries := []string{}
		postalCodes := []string{}
		description := ""
		amenities := []string{}
		uniqueImages := map[string]schema.Image{}
		uniqueBookingConditions := map[string]bool{}

		for _, hotel := range hotels {
			if hotel.Name != "" {
				names = append(names, hotel.Name)
			}
			if hotel.Latitude != nil {
				latitudes = append(latitudes, *hotel.Latitude)
			}
			if hotel.Longitude != nil {
				longitudes = append(longitudes, *hotel.Longitude)
			}
			if hotel.City != "" {
				cities = append(cities, hotel.City)
			}
			if hotel.Country != "" {
				countries = append(countries, hotel.Country)
			}
			if hotel.PostalCode != "" {
				postalCodes = append(postalCodes, hotel.PostalCode)
			}
			amenities = append(amenities, hotel.Amenities...)

			// Find the longest address
			if len(hotel.Address) > len(longestAddress) {
				longestAddress = hotel.Address
			}

			// Combine descriptions
			description += hotel.Description

			// Combine images
			for _, image := range hotel.Images {
				uniqueImages[image.URL] = image
			}

			// Combine booking conditions
			for _, condition := range hotel.BookingConditions {
				uniqueBookingConditions[condition] = true
			}
		}

		// Merge image by URL
		mergedImages := []schema.Image{}
		for _, image := range uniqueImages {
			mergedImages = append(mergedImages, image)
		}

		// Merge booking conditions by uniqueness
		mergedBookingConditions := []string{}
		for condition := range uniqueBookingConditions {
			mergedBookingConditions = append(mergedBookingConditions, condition)
		}

		// Merge hotels with the same ID
		mergedHotel := &ent.Hotel{
			ID:                hotels[0].ID,
			DestinationID:     hotels[0].DestinationID,
			Name:              utils.FindMostAverageString(names),
			Address:           longestAddress,
			City:              utils.FindMostAverageString(cities),
			Country:           utils.FindMostAverageString(countries),
			PostalCode:        utils.FindMostAverageString(postalCodes),
			Description:       description,
			Amenities:         utils.GetUniqueStrings(amenities),
			Images:            mergedImages,
			BookingConditions: mergedBookingConditions,
		}

		// Calculate average latitude
		if len(latitudes) > 0 {
			average := utils.CalculateAverage(latitudes)
			mergedHotel.Latitude = &average
		}

		// Calculate average longitude
		if len(longitudes) > 0 {
			average := utils.CalculateAverage(longitudes)
			mergedHotel.Longitude = &average
		}

		unifiedHotels = append(unifiedHotels, mergedHotel)
	}

	return unifiedHotels, nil
}
