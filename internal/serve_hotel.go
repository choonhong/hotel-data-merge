package internal

import (
	"net/http"

	"github.com/choonhong/hotel-data-merge/restapi"
	"github.com/choonhong/hotel-data-merge/utils"
)

func (s *HotelService) GetHotels(w http.ResponseWriter, r *http.Request, params restapi.GetHotelsParams) {
	hotels, err := s.HotelRepo.GetHotels(r.Context(), params.Ids, params.Destination)
	if err != nil {
		utils.ResponseHelper(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseHelper(w, hotels, http.StatusOK)
}
