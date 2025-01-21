package adapter

import (
	"context"

	"github.com/choonhong/hotel-data-merge/ent"
	"github.com/choonhong/hotel-data-merge/ent/hotel"
)

type HotelRepository struct {
	*ent.Client
}

func (h *HotelRepository) Save(ctx context.Context, hotels *ent.Hotel) error {
	create := h.Client.Hotel.Create().
		SetID(hotels.ID).
		SetDestinationID(hotels.DestinationID).
		SetName(hotels.Name).
		SetAddress(hotels.Address).
		SetCity(hotels.City).
		SetCountry(hotels.Country).
		SetPostalCode(hotels.PostalCode).
		SetDescription(hotels.Description).
		SetAmenities(hotels.Amenities).
		SetImages(hotels.Images).
		SetBookingConditions(hotels.BookingConditions)

	if hotels.Latitude != nil {
		create.SetLatitude(*hotels.Latitude)
	}

	if hotels.Longitude != nil {
		create.SetLongitude(*hotels.Longitude)
	}

	return create.OnConflict().UpdateNewValues().Exec(ctx)
}

func (h *HotelRepository) GetHotels(ctx context.Context, ids []string, destinationID int) ([]*ent.Hotel, error) {
	query := h.Client.Hotel.Query()

	if len(ids) != 0 {
		query = query.Where(hotel.IDIn(ids...))
	}

	if destinationID != 0 {
		query = query.Where(hotel.DestinationIDEQ(destinationID))
	}

	return query.All(ctx)
}
