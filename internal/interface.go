package internal

import (
	"context"

	"github.com/choonhong/hotel-data-merge/ent"
)

type HotelRepository interface {
	Save(ctx context.Context, hotels *ent.Hotel) error
	GetHotels(ctx context.Context, ids *[]string, destinationID *int) ([]*ent.Hotel, error)
}

type Provider interface {
	Name() string
	FetchAll(context.Context) ([]*ent.Hotel, error)
}

type Cache interface {
	Get(v interface{}) ([]*ent.Hotel, error)
	Set(v interface{}, hotels []*ent.Hotel) error
	Clear() error
}
