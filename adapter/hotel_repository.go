package adapter

import (
	"github.com/choonhong/hotel-data-merge/model"
	"gorm.io/gorm"
)

type HotelRepository struct {
	*gorm.DB
}

func (h *HotelRepository) SaveAll(hotels []*model.Hotel) error {
	return h.Create(hotels).Error
}

func (h *HotelRepository) GetAll() ([]*model.Hotel, error) {
	var hotels []*model.Hotel
	err := h.Find(&hotels).Error

	return hotels, err
}

func (h *HotelRepository) GetByIDs(ids []string) ([]*model.Hotel, error) {
	var hotels []*model.Hotel
	err := h.Where("id IN ?", ids).Find(&hotels).Error

	return hotels, err
}

func (h *HotelRepository) GetByDestinationID(destinationID int) ([]*model.Hotel, error) {
	var hotels []*model.Hotel
	err := h.Where("destination_id = ?", destinationID).Find(&hotels).Error

	return hotels, err
}

func (h *HotelRepository) GetByIDsAndDestinationID(ids []string, destinationID int) ([]*model.Hotel, error) {
	var hotels []*model.Hotel
	err := h.Where("id IN ? AND destination_id = ?", ids, destinationID).Find(&hotels).Error

	return hotels, err
}
