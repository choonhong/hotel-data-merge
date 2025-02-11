package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/choonhong/hotel-data-merge/ent"
	"github.com/choonhong/hotel-data-merge/utils"
)

type Acme struct {
	URL               string
	FetchDataFunction func(ctx context.Context, url string) (io.ReadCloser, error)
}

func (a *Acme) Name() string {
	return "Acme"
}

// see testdata/acme.json
type AcmeData struct {
	ID            string          `json:"Id"`
	DestinationID int             `json:"DestinationId"`
	Name          string          `json:"Name"`
	Latitude      Float64OrString `json:"Latitude"`
	Longitude     Float64OrString `json:"Longitude"`
	Address       string          `json:"Address"`
	City          string          `json:"City"`
	Country       string          `json:"Country"`
	PostalCode    string          `json:"PostalCode"`
	Description   string          `json:"Description"`
	Facilities    []string        `json:"Facilities"`
}

func (d *AcmeData) ToHotel() *ent.Hotel {
	return &ent.Hotel{
		ID:            d.ID,
		DestinationID: d.DestinationID,
		Name:          strings.TrimSpace(d.Name),
		Latitude:      d.Latitude.Value,
		Longitude:     d.Longitude.Value,
		Address:       strings.TrimSpace(d.Address),
		City:          strings.TrimSpace(d.City),
		Country:       strings.TrimSpace(d.Country),
		PostalCode:    strings.TrimSpace(d.PostalCode),
		Description:   strings.TrimSpace(d.Description),
		Amenities:     utils.TrimSpaceInList(d.Facilities),
	}
}

// FetchAll fetches all hotels from Acme API.
func (a *Acme) FetchAll(ctx context.Context) ([]*ent.Hotel, error) {
	// Call data from Acme API
	body, err := a.FetchDataFunction(ctx, a.URL)
	if err != nil {
		return nil, fmt.Errorf("fetchDataFromURL: %w", err)
	}
	defer body.Close()

	// Decode the response body
	var data []*AcmeData
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return nil, fmt.Errorf("Decode: %w", err)
	}

	// Parse data to internal Hotel model
	var hotels []*ent.Hotel
	for _, d := range data {
		hotels = append(hotels, d.ToHotel())
	}

	return hotels, nil
}
