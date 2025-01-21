package adapter

import (
	"context"

	"github.com/choonhong/hotel-data-merge/domain"
	"github.com/choonhong/hotel-data-merge/ent"
	"github.com/choonhong/hotel-data-merge/ent/hotel"
)

type HotelRepository struct {
	*ent.Client
}

// Save saves the unified hotel data to the database
func (d *HotelRepository) Save(ctx context.Context, data *domain.Hotel) error {
	_, err := d.Hotel.Create().
		SetID(data.ID).
		SetDestinationID(data.DestinationID).
		SetName(data.Name).
		SetLatitude(data.Latitude).
		SetLongitude(data.Longitude).
		SetAddress(data.Address).
		SetCity(data.City).
		SetCountry(data.Country).
		SetDetails(data.Details).
		SetAmenities(data.Amenities).
		SetImages(data.Images).
		Save(ctx)

	return err
}

// GetByIDs retrieves the hotels data by IDs
func (d *HotelRepository) GetByIDs(ctx context.Context, ids []string) ([]*domain.Hotel, error) {
	hotels, err := d.Hotel.Query().
		Where(hotel.IDIn(ids...)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return toDomainHotels(hotels), nil
}

// GetByDestinationID retrieves the hotels data by destination ID
func (d *HotelRepository) GetByDestinationID(ctx context.Context, destinationID string) ([]*domain.Hotel, error) {
	hotels, err := d.Hotel.Query().
		Where(hotel.DestinationID(destinationID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return toDomainHotels(hotels), nil
}

// parse ent.Hotel to domain.Hotel
func toDomainHotels(hotels []*ent.Hotel) []*domain.Hotel {
	var result []*domain.Hotel

	for _, h := range hotels {
		result = append(result, &domain.Hotel{
			ID:            h.ID,
			DestinationID: h.DestinationID,
			Name:          h.Name,
			Latitude:      h.Latitude,
			Longitude:     h.Longitude,
			Address:       h.Address,
			City:          h.City,
			Country:       h.Country,
			Details:       h.Details,
			Amenities:     h.Amenities,
			Images:        h.Images,
		})
	}

	return result
}
