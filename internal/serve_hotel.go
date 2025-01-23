package internal

import (
	"log"
	"net/http"

	"github.com/choonhong/hotel-data-merge/restapi"
	"github.com/choonhong/hotel-data-merge/utils"
)

func (s *HotelService) GetHotels(w http.ResponseWriter, r *http.Request, params restapi.GetHotelsParams) {
	log.Println("Get hotels")

	// Check cache first
	cachedHotels, err := s.Cache.Get(params)
	if err == nil {
		log.Println("Get hotels from cache")
		utils.ResponseHelper(w, cachedHotels, http.StatusOK)
		return
	}

	// Fetch hotels from the database
	hotels, err := s.HotelRepo.GetHotels(r.Context(), params.Ids, params.Destination)
	if err != nil {
		log.Println("Error fetching hotels:", err)
		utils.ResponseHelper(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save to cache
	if err := s.Cache.Set(params, hotels); err != nil {
		log.Println("Error saving to cache:", err)
	}

	utils.ResponseHelper(w, hotels, http.StatusOK)
}
