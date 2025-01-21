package internal

import (
	"context"

	"github.com/choonhong/hotel-data-merge/model"
)

type HotelRepo interface {
	SaveAll(hotels []*model.Hotel) error
	GetAll() ([]*model.Hotel, error)
	GetByIDs(ids []string) ([]*model.Hotel, error)
	GetByDestinationID(destinationID int) ([]*model.Hotel, error)
	GetByIDsAndDestinationID(ids []string, destinationID int) ([]*model.Hotel, error)
}

type Provider interface {
	Name() string
	FetchAll(context.Context) ([]*model.Hotel, error)
}
