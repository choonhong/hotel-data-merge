package adapter_test

import (
	"context"
	"testing"

	"github.com/choonhong/hotel-data-merge/adapter"
	"github.com/choonhong/hotel-data-merge/database"
	"github.com/choonhong/hotel-data-merge/ent"
	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	db, err := database.Connect()
	require.NoError(t, err)
	defer db.Close()

	subject := adapter.HotelRepository{Client: db}
	hotel := ent.Hotel{ID: "1", Name: "Hotel A", DestinationID: 1}

	err = subject.Save(context.Background(), &hotel)
	require.NoError(t, err)

	hotel.Name = "Hotel B"
	err = subject.Save(context.Background(), &hotel)
	require.NoError(t, err)

	hotels, err := db.Hotel.Query().All(context.Background())
	require.NoError(t, err)
	require.Len(t, hotels, 1)
	require.Equal(t, "Hotel B", hotels[0].Name)
}

func TestGetHotels(t *testing.T) {
	db, err := database.Connect()
	require.NoError(t, err)
	defer db.Close()

	subject := adapter.HotelRepository{Client: db}

	hotel := ent.Hotel{ID: "1", Name: "Hotel A", DestinationID: 1}
	err = subject.Save(context.Background(), &hotel)
	require.NoError(t, err)

	hotel2 := ent.Hotel{ID: "2", Name: "Hotel B", DestinationID: 2}
	err = subject.Save(context.Background(), &hotel2)
	require.NoError(t, err)

	hotel3 := ent.Hotel{ID: "3", Name: "Hotel C", DestinationID: 1}
	err = subject.Save(context.Background(), &hotel3)
	require.NoError(t, err)

	hotels, err := subject.GetHotels(context.Background(), nil, nil)
	require.NoError(t, err)
	require.Len(t, hotels, 3)
	require.Equal(t, "Hotel A", hotels[0].Name)

	hotels, err = subject.GetHotels(context.Background(), &[]string{"1"}, nil)
	require.NoError(t, err)
	require.Len(t, hotels, 1)
	require.Equal(t, "Hotel A", hotels[0].Name)

	hotels, err = subject.GetHotels(context.Background(), nil, &hotel.DestinationID)
	require.NoError(t, err)
	require.Len(t, hotels, 2)
	require.Equal(t, "Hotel A", hotels[0].Name)
	require.Equal(t, "Hotel C", hotels[1].Name)

	hotels, err = subject.GetHotels(context.Background(), &[]string{"1", "2"}, nil)
	require.NoError(t, err)
	require.Len(t, hotels, 2)
	require.Equal(t, "Hotel A", hotels[0].Name)
	require.Equal(t, "Hotel B", hotels[1].Name)

	hotels, err = subject.GetHotels(context.Background(), &[]string{"1", "2"}, &hotel.DestinationID)
	require.NoError(t, err)
	require.Len(t, hotels, 1)
	require.Equal(t, "Hotel A", hotels[0].Name)
}
