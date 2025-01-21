package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/choonhong/hotel-data-merge/model"
)

type Acme struct {
	URL string
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

func (a *Acme) FetchAll(ctx context.Context) ([]*model.Hotel, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Do: %w", err)
	}
	defer resp.Body.Close()

	var data []*AcmeData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("Decode: %w", err)
	}

	var hotels []*model.Hotel
	for _, d := range data {
		hotels = append(hotels, &model.Hotel{
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
			Amenities:     trimSpaceInList(d.Facilities),
		})
	}

	return hotels, nil
}
