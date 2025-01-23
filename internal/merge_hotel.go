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
	HotelRepo HotelRepository
	Providers []Provider
	Cache     Cache
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

	hotelMap := groupHotelsByID(hotels)
	unifiedHotels := map[string]*ent.Hotel{}

	// Merge hotels with the same ID
	for hotelID, hotels := range hotelMap {
		unifiedHotels[hotelID] = mergeHotelGroup(hotels)
	}

	return unifiedHotels
}

func groupHotelsByID(hotels []*ent.Hotel) map[string][]*ent.Hotel {
	hotelMap := map[string][]*ent.Hotel{}
	for _, hotel := range hotels {
		hotelMap[hotel.ID] = append(hotelMap[hotel.ID], hotel)
	}
	return hotelMap
}

func mergeHotelGroup(hotels []*ent.Hotel) *ent.Hotel {
	if len(hotels) == 1 {
		return hotels[0]
	}

	hotelData := collectHotelData(hotels)

	mergedHotel := &ent.Hotel{
		ID:                hotels[0].ID,
		DestinationID:     hotels[0].DestinationID,
		Name:              utils.FindMostRepresentativeString(hotelData.names),
		Address:           findLongestAddress(hotels),
		City:              utils.FindMostRepresentativeString(hotelData.cities),
		Country:           utils.FindMostRepresentativeString(hotelData.countries),
		PostalCode:        utils.FindMostRepresentativeString(hotelData.postalCodes),
		Description:       mergeDescriptions(hotels),
		Amenities:         utils.RemoveSimilarAndSubstrings(hotelData.amenities),
		Images:            mergeImages(hotelData.uniqueImages),
		BookingConditions: mergeBookingConditions(hotelData.uniqueBookingConditions),
	}

	if len(hotelData.latitudes) > 0 {
		mergedHotel.Latitude = utils.CalculateAverage(hotelData.latitudes)
	}

	if len(hotelData.longitudes) > 0 {
		mergedHotel.Longitude = utils.CalculateAverage(hotelData.longitudes)
	}

	log.Println("Merged hotel", mergedHotel.ID)
	return mergedHotel
}

// hotelData holds the collected data of the same hotel from different providers.
type hotelData struct {
	names                   []string
	latitudes               []float64
	longitudes              []float64
	cities                  []string
	countries               []string
	postalCodes             []string
	amenities               []string
	uniqueImages            map[string]schema.Image
	uniqueBookingConditions map[string]bool
}

func collectHotelData(hotels []*ent.Hotel) hotelData {
	hotelData := hotelData{
		uniqueImages:            map[string]schema.Image{},
		uniqueBookingConditions: map[string]bool{},
	}

	for _, hotel := range hotels {
		if hotel.Name != "" {
			hotelData.names = append(hotelData.names, hotel.Name)
		}
		if hotel.Latitude != nil {
			hotelData.latitudes = append(hotelData.latitudes, *hotel.Latitude)
		}
		if hotel.Longitude != nil {
			hotelData.longitudes = append(hotelData.longitudes, *hotel.Longitude)
		}
		if hotel.City != "" {
			hotelData.cities = append(hotelData.cities, hotel.City)
		}
		if hotel.Country != "" {
			hotelData.countries = append(hotelData.countries, hotel.Country)
		}
		if hotel.PostalCode != "" {
			hotelData.postalCodes = append(hotelData.postalCodes, hotel.PostalCode)
		}
		hotelData.amenities = append(hotelData.amenities, hotel.Amenities...)
		for _, image := range hotel.Images {
			hotelData.uniqueImages[image.URL] = image
		}
		for _, condition := range hotel.BookingConditions {
			hotelData.uniqueBookingConditions[condition] = true
		}
	}

	return hotelData
}

func findLongestAddress(hotels []*ent.Hotel) string {
	longestAddress := ""
	for _, hotel := range hotels {
		if len(hotel.Address) > len(longestAddress) {
			longestAddress = hotel.Address
		}
	}
	return longestAddress
}

func mergeDescriptions(hotels []*ent.Hotel) string {
	description := ""
	for _, hotel := range hotels {
		description += hotel.Description
	}
	return description
}

func mergeImages(uniqueImages map[string]schema.Image) []schema.Image {
	mergedImages := []schema.Image{}
	for _, image := range uniqueImages {
		mergedImages = append(mergedImages, image)
	}
	return mergedImages
}

func mergeBookingConditions(uniqueBookingConditions map[string]bool) []string {
	mergedBookingConditions := []string{}
	for condition := range uniqueBookingConditions {
		mergedBookingConditions = append(mergedBookingConditions, condition)
	}
	return mergedBookingConditions
}
