package internal

import (
	"github.com/choonhong/hotel-data-merge/model"
)

func (s *HotelService) GetHotels(id []string, destinationID int) ([]*model.Hotel, error) {
	if len(id) == 0 && destinationID == 0 {
		return s.HotelRepo.GetAll()
	}

	if destinationID == 0 {
		return s.HotelRepo.GetByIDs(id)
	}

	if len(id) == 0 {
		return s.HotelRepo.GetByDestinationID(destinationID)
	}

	return s.HotelRepo.GetByIDsAndDestinationID(id, destinationID)
}
