package internal

import (
	"context"
	"fmt"

	"github.com/agext/levenshtein"
	"github.com/choonhong/hotel-data-merge/model"
	"github.com/davecgh/go-spew/spew"
)

type HotelService struct {
	HotelRepo HotelRepo
	Providers []Provider
}

func (s *HotelService) FetchAndMergeHotels(ctx context.Context) ([]*model.Hotel, error) {
	hotelMap, err := s.fetchHotels(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetchHotels: %w", err)
	}

	hotels, err := s.mergeHotels(hotelMap)
	if err != nil {
		return nil, fmt.Errorf("mergeHotels: %w", err)
	}

	return hotels, nil
}

func (s *HotelService) fetchHotels(ctx context.Context) (map[string][]*model.Hotel, error) {
	hotelMap := map[string][]*model.Hotel{}

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
				hotelMap[hotel.ID] = []*model.Hotel{}
			}

			hotelMap[hotel.ID] = append(hotelMap[hotel.ID], hotel)
		}
	}

	return hotelMap, nil
}

func (s *HotelService) mergeHotels(hotelMap map[string][]*model.Hotel) ([]*model.Hotel, error) {
	fmt.Println("Merging hotels...")

	unifiedHotels := []*model.Hotel{}

	// Merge hotels with the same ID
	for _, hotels := range hotelMap {
		if len(hotels) == 1 {
			unifiedHotels = append(unifiedHotels, hotels[0])
			continue
		}

		names := []string{}
		latitudes := []float64{}
		longitudes := []float64{}
		longestAddress := ""
		cities := []string{}
		countries := []string{}
		postalCodes := []string{}
		description := ""
		uniqueAmenities := map[string]bool{}
		uniqueImages := map[string]model.Image{}
		uniqueBookingConditions := map[string]bool{}

		for _, hotel := range hotels {
			if hotel.Name != "" {
				names = append(names, hotel.Name)
			}
			if hotel.Latitude != nil {
				latitudes = append(latitudes, *hotel.Latitude)
			}
			if hotel.Latitude != nil {
				longitudes = append(longitudes, *hotel.Latitude)
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

			// Find the longest address
			if len(hotel.Address) > len(longestAddress) {
				longestAddress = hotel.Address
			}

			// Combine descriptions
			description += hotel.Description

			// Combine amenities
			for _, amenity := range hotel.Amenities {
				uniqueAmenities[amenity] = true
			}

			// Combine images
			for _, image := range hotel.Images {
				uniqueImages[image.URL] = image
			}

			// Combine booking conditions
			for _, condition := range hotel.BookingConditions {
				uniqueBookingConditions[condition] = true
			}
		}

		mergedAmenities := []string{}
		for amenity := range uniqueAmenities {
			mergedAmenities = append(mergedAmenities, amenity)
		}

		mergedImages := []model.Image{}
		for _, image := range uniqueImages {
			mergedImages = append(mergedImages, image)
		}

		mergedBookingConditions := []string{}
		for condition := range uniqueBookingConditions {
			mergedBookingConditions = append(mergedBookingConditions, condition)
		}

		// Merge hotels with the same ID
		mergedHotel := &model.Hotel{
			ID:                hotels[0].ID,
			DestinationID:     hotels[0].DestinationID,
			Name:              findMostAverageString(names),
			Address:           longestAddress,
			City:              findMostAverageString(cities),
			Country:           findMostAverageString(countries),
			PostalCode:        findMostAverageString(postalCodes),
			Description:       description,
			Amenities:         mergedAmenities,
			Images:            mergedImages,
			BookingConditions: mergedBookingConditions,
		}

		if len(latitudes) > 0 {
			average := calculateAverage(latitudes)
			mergedHotel.Latitude = &average
		}

		if len(longitudes) > 0 {
			average := calculateAverage(longitudes)
			mergedHotel.Longitude = &average
		}

		unifiedHotels = append(unifiedHotels, mergedHotel)
	}

	spew.Dump(unifiedHotels)

	// Save merged hotels to the hotel repository
	if err := s.HotelRepo.SaveAll(unifiedHotels); err != nil {
		return nil, fmt.Errorf("SaveAll: %w", err)
	}

	return unifiedHotels, nil
}

// findMostAverageString combines multiple strings into one, returning the most average string
func findMostAverageString(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	minTotalDistance := len(strs[0]) * len(strs)
	averageName := ""

	for _, str1 := range strs {
		totalDistance := 0

		for _, str2 := range strs {
			if str1 != str2 {
				totalDistance += levenshtein.Distance(str1, str2, nil)
			}
		}

		// no other name to compare
		if totalDistance == 0 {
			return str1
		}

		if totalDistance < minTotalDistance {
			minTotalDistance = totalDistance
			averageName = str1
		}
	}

	return averageName
}

func calculateAverage(list []float64) float64 {
	total := 0.0
	for _, num := range list {
		total += num
	}
	return total / float64(len(list))
}
