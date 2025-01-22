package internal

import (
	"log"
	"net/http"

	"github.com/choonhong/hotel-data-merge/restapi"
	"github.com/choonhong/hotel-data-merge/utils"
)

func (s *HotelService) GetHotels(w http.ResponseWriter, r *http.Request, params restapi.GetHotelsParams) {
	log.Println("GetHotels", params)

	hotels, err := s.HotelRepo.GetHotels(r.Context(), params.Ids, params.Destination)
	if err != nil {
		log.Println("Error fetching hotels:", err)
		utils.ResponseHelper(w, err.Error(), http.StatusInternalServerError)

		return
	}

	utils.ResponseHelper(w, hotels, http.StatusOK)
}
