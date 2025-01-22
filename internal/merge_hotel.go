package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/choonhong/hotel-data-merge/ent"
	"github.com/choonhong/hotel-data-merge/ent/schema"
	"github.com/choonhong/hotel-data-merge/utils"
)

type HotelService struct {
	HotelRepo HotelRepo
	Providers []Provider
}

// FetchAndMergeHotels fetches hotels from all providers and merges them into a single hotel list.
func (s *HotelService) FetchAndMergeHotels(ctx context.Context) (map[string]*ent.Hotel, error) {
	// Fetch hotels from all providers
	allHotels, err := s.fetchHotels(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetchHotels: %w", err)
	}

	// Merge hotel of the same ID
	mergedHotels := MergeHotels(allHotels)

	log.Println("Saving", len(mergedHotels), "hotels...")

	// Save merged hotels
	for _, hotel := range mergedHotels {
		if err := s.HotelRepo.Save(context.Background(), hotel); err != nil {
			return nil, fmt.Errorf("Save: %w", err)
		}
	}

	log.Println("Fetch and merge hotels completed.")

	return mergedHotels, nil
}

// fetchHotels fetches hotels from all providers and groups them by ID.
func (s *HotelService) fetchHotels(ctx context.Context) ([]*ent.Hotel, error) {
	allHotels := []*ent.Hotel{}

	// Fetch hotels from all providers
	for _, provider := range s.Providers {
		log.Println("Fetching hotels from", provider.Name())

		hotels, err := provider.FetchAll(ctx)
		if err != nil {
			// Log error and continue to the next provider
			log.Println("Error fetching hotels from", provider.Name(), ":", err)
			continue
		}

		allHotels = append(allHotels, hotels...)
	}

	return allHotels, nil
}

// MergeHotels merges hotels with the same ID into a single hotel list.
func MergeHotels(hotels []*ent.Hotel) map[string]*ent.Hotel {
	log.Println("Merging hotels...")

	hotelMap := map[string][]*ent.Hotel{}
	for _, hotel := range hotels {
		if _, ok := hotelMap[hotel.ID]; !ok {
			hotelMap[hotel.ID] = []*ent.Hotel{}
		}
		hotelMap[hotel.ID] = append(hotelMap[hotel.ID], hotel)
	}

	unifiedHotels := map[string]*ent.Hotel{}

	// Merge hotels with the same ID
	for _, hotels := range hotelMap {
		if len(hotels) == 1 {
			unifiedHotels[hotels[0].ID] = hotels[0]
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
			Name:              utils.FindMostRepresentativeString(names),
			Address:           longestAddress,
			City:              utils.FindMostRepresentativeString(cities),
			Country:           utils.FindMostRepresentativeString(countries),
			PostalCode:        utils.FindMostRepresentativeString(postalCodes),
			Description:       description,
			Amenities:         utils.RemoveSimilarAndSubstrings(amenities),
			Images:            mergedImages,
			BookingConditions: mergedBookingConditions,
		}

		// Calculate average latitude
		if len(latitudes) > 0 {
			mergedHotel.Latitude = utils.CalculateAverage(latitudes)
		}

		// Calculate average longitude
		if len(longitudes) > 0 {
			mergedHotel.Longitude = utils.CalculateAverage(longitudes)
		}

		unifiedHotels[mergedHotel.ID] = mergedHotel

		log.Println("Merged hotel", mergedHotel.ID)
	}

	return unifiedHotels
}
