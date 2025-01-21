package internal

import (
	"context"

	"github.com/choonhong/hotel-data-merge/ent"
)

func (s *HotelService) GetHotels(ctx context.Context, id []string, destinationID int) ([]*ent.Hotel, error) {
	return s.HotelRepo.GetHotels(ctx, id, destinationID)
}
