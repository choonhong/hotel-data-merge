package adapter

import (
	"context"

	"github.com/choonhong/hotel-data-merge/domain"
	"github.com/choonhong/hotel-data-merge/ent"
)

type DataRepository struct {
	*ent.Client
}

// Save saves the raw hotel data and the provider to the database
func (d *DataRepository) Save(ctx context.Context, data *domain.Hotel, provider string) error {
	_, err := d.Data.Create().
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
		SetProvider(provider).
		Save(ctx)

	return err
}
