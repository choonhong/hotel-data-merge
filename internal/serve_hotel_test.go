package internal_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/choonhong/hotel-data-merge/ent"
	"github.com/choonhong/hotel-data-merge/internal"
	"github.com/choonhong/hotel-data-merge/internal/mocks"
	"github.com/choonhong/hotel-data-merge/restapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHotels(t *testing.T) {
	t.Run("cache hit", func(t *testing.T) {
		// Create a new HotelService with mocks
		subject, mocks := newHotelServiceMocks()

		// Create a new request
		r := httptest.NewRequest(http.MethodGet, "/hotels", nil)
		w := httptest.NewRecorder()

		// Create a new GetHotelsParams
		params := restapi.GetHotelsParams{Ids: &[]string{"iJhz"}}
		hotels := []*ent.Hotel{{ID: "iJhz", Name: "Hotel A", DestinationID: 5432}}

		// Mocks function response
		mocks.cache.EXPECT().Get(params).Return(hotels, nil)

		// Call the GetHotels function
		subject.GetHotels(w, r, params)

		// Check the response
		assert.Equal(t, http.StatusOK, w.Code)
		responseHotel := []*ent.Hotel{}
		err := json.Unmarshal(w.Body.Bytes(), &responseHotel)
		require.NoError(t, err)
		assert.Equal(t, hotels, responseHotel)
	})

	t.Run("database error", func(t *testing.T) {
		// Create a new HotelService with mocks
		subject, mocks := newHotelServiceMocks()

		// Create a new request
		r := httptest.NewRequest(http.MethodGet, "/hotels", nil)
		w := httptest.NewRecorder()

		// Create a new GetHotelsParams
		params := restapi.GetHotelsParams{Ids: &[]string{"iJhz"}}

		// Mocks function response
		mocks.cache.EXPECT().Get(params).Return(nil, assert.AnError)
		mocks.hotelRepo.EXPECT().GetHotels(r.Context(), &[]string{"iJhz"}, (*int)(nil)).Return(nil, assert.AnError)

		// Call the GetHotels function
		subject.GetHotels(w, r, params)

		// Check the response status code
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("get from database", func(t *testing.T) {
		// Create a new HotelService with mocks
		subject, mocks := newHotelServiceMocks()

		// Create a new request
		r := httptest.NewRequest(http.MethodGet, "/hotels", nil)
		w := httptest.NewRecorder()

		// Create a new GetHotelsParams
		params := restapi.GetHotelsParams{Ids: &[]string{"iJhz"}}
		hotels := []*ent.Hotel{{ID: "iJhz", Name: "Hotel A", DestinationID: 5432}}

		// Mocks function response
		mocks.cache.EXPECT().Get(params).Return(nil, assert.AnError)
		mocks.hotelRepo.EXPECT().GetHotels(r.Context(), &[]string{"iJhz"}, (*int)(nil)).Return(hotels, nil)
		mocks.cache.EXPECT().Set(params, hotels).Return(nil)

		// Call the GetHotels function
		subject.GetHotels(w, r, params)

		// Check the response
		assert.Equal(t, http.StatusOK, w.Code)
		responseHotel := []*ent.Hotel{}
		err := json.Unmarshal(w.Body.Bytes(), &responseHotel)
		require.NoError(t, err)
		assert.Equal(t, hotels, responseHotel)
	})

	t.Run("save to cache error", func(t *testing.T) {
		// Create a new HotelService with mocks
		subject, mocks := newHotelServiceMocks()

		// Create a new request
		r := httptest.NewRequest(http.MethodGet, "/hotels", nil)
		w := httptest.NewRecorder()

		// Create a new GetHotelsParams
		params := restapi.GetHotelsParams{Ids: &[]string{"iJhz"}}
		hotels := []*ent.Hotel{{ID: "iJhz", Name: "Hotel A", DestinationID: 5432}}

		// Mocks function response
		mocks.cache.EXPECT().Get(params).Return(nil, assert.AnError)
		mocks.hotelRepo.EXPECT().GetHotels(r.Context(), &[]string{"iJhz"}, (*int)(nil)).Return(hotels, nil)
		mocks.cache.EXPECT().Set(params, hotels).Return(assert.AnError)

		// Call the GetHotels function
		subject.GetHotels(w, r, params)

		// Check the response
		assert.Equal(t, http.StatusOK, w.Code)
		responseHotel := []*ent.Hotel{}
		err := json.Unmarshal(w.Body.Bytes(), &responseHotel)
		require.NoError(t, err)
		assert.Equal(t, hotels, responseHotel)
	})
}

type hotelServiceMocks struct {
	cache     *mocks.Cache
	hotelRepo *mocks.HotelRepository
	provoder  *mocks.Provider
}

func newHotelServiceMocks() (internal.HotelService, hotelServiceMocks) {
	mocks := hotelServiceMocks{
		cache:     &mocks.Cache{},
		hotelRepo: &mocks.HotelRepository{},
		provoder:  &mocks.Provider{},
	}

	return internal.HotelService{
		HotelRepo: mocks.hotelRepo,
		Cache:     mocks.cache,
		Providers: []internal.Provider{mocks.provoder},
	}, mocks
}
