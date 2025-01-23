package adapter_test

import (
	"testing"

	"github.com/choonhong/hotel-data-merge/adapter"
	"github.com/choonhong/hotel-data-merge/ent"
	"github.com/choonhong/hotel-data-merge/restapi"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("Get non-existent", func(t *testing.T) {
		cache := adapter.NewCache()
		params := restapi.GetHotelsParams{Ids: &[]string{"1"}}
		cachedHotels, err := cache.Get(params)
		require.Error(t, err)
		require.Nil(t, cachedHotels)
	})

	t.Run("Empty query", func(t *testing.T) {
		cache := adapter.NewCache()
		params := restapi.GetHotelsParams{}
		hotels := []*ent.Hotel{
			{ID: "1", Name: "Hotel A", DestinationID: 1},
			{ID: "2", Name: "Hotel B", DestinationID: 2},
		}

		err := cache.Set(params, hotels)
		require.NoError(t, err)

		cachedHotels, err := cache.Get(params)
		require.NoError(t, err)
		require.Equal(t, hotels, cachedHotels)
	})

	t.Run("Query with IDs", func(t *testing.T) {
		cache := adapter.NewCache()
		params := restapi.GetHotelsParams{Ids: &[]string{"1"}}
		hotels := []*ent.Hotel{
			{ID: "1", Name: "Hotel A", DestinationID: 1},
		}

		err := cache.Set(params, hotels)
		require.NoError(t, err)

		cachedHotels, err := cache.Get(params)
		require.NoError(t, err)
		require.Equal(t, hotels, cachedHotels)
	})

	t.Run("Query with destination", func(t *testing.T) {
		cache := adapter.NewCache()
		destination := 1
		params := restapi.GetHotelsParams{Destination: &destination}
		hotels := []*ent.Hotel{
			{ID: "1", Name: "Hotel A", DestinationID: 1},
		}

		err := cache.Set(params, hotels)
		require.NoError(t, err)

		cachedHotels, err := cache.Get(params)
		require.NoError(t, err)
		require.Equal(t, hotels, cachedHotels)
	})

	t.Run("Store multiple queries", func(t *testing.T) {
		cache := adapter.NewCache()
		params1 := restapi.GetHotelsParams{Ids: &[]string{"1"}}
		hotels1 := []*ent.Hotel{
			{ID: "1", Name: "Hotel A", DestinationID: 1},
		}
		params2 := restapi.GetHotelsParams{Ids: &[]string{"2"}}
		hotels2 := []*ent.Hotel{
			{ID: "2", Name: "Hotel B", DestinationID: 1},
		}

		err := cache.Set(params1, hotels1)
		require.NoError(t, err)
		err = cache.Set(params2, hotels2)
		require.NoError(t, err)

		cachedHotels1, err := cache.Get(params1)
		require.NoError(t, err)
		require.Equal(t, hotels1, cachedHotels1)

		cachedHotels2, err := cache.Get(params2)
		require.NoError(t, err)
		require.Equal(t, hotels2, cachedHotels2)
	})

	t.Run("Clear", func(t *testing.T) {
		cache := adapter.NewCache()
		params := restapi.GetHotelsParams{Ids: &[]string{"1"}}
		hotels := []*ent.Hotel{
			{ID: "1", Name: "Hotel A", DestinationID: 1},
		}

		err := cache.Set(params, hotels)
		require.NoError(t, err)

		err = cache.Clear()
		require.NoError(t, err)

		cachedHotels, err := cache.Get(params)
		require.Error(t, err)
		require.Nil(t, cachedHotels)
	})
}
