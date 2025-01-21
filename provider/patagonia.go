package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/choonhong/hotel-data-merge/ent"
	"github.com/choonhong/hotel-data-merge/ent/schema"
)

type Patagonia struct {
	URL string
}

func (p *Patagonia) Name() string {
	return "Patagonia"
}

// see testdata/patagonia.json
type PatagoniaData struct {
	ID            string                      `json:"id"`
	DestinationID int                         `json:"destination"`
	Name          string                      `json:"name"`
	Latitude      Float64OrString             `json:"lat"`
	Longitude     Float64OrString             `json:"lng"`
	Address       string                      `json:"address"`
	Info          string                      `json:"info"`
	Amenities     []string                    `json:"amenities"`
	Images        map[string][]PatagoniaImage `json:"images"`
}

type PatagoniaImage struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

// FetchAll fetches all hotels from Patagonia API.
func (p *Patagonia) FetchAll(ctx context.Context) ([]*ent.Hotel, error) {
	// Call Patagonia API
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.URL, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Do: %w", err)
	}
	defer resp.Body.Close()

	// Decode the response body
	var data []*PatagoniaData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("Decode: %w", err)
	}

	// Parse data to internal Hotel model
	var hotels []*ent.Hotel
	for _, d := range data {
		hotel := &ent.Hotel{
			ID:            d.ID,
			DestinationID: d.DestinationID,
			Name:          strings.TrimSpace(d.Name),
			Latitude:      d.Latitude.Value,
			Longitude:     d.Longitude.Value,
			Address:       strings.TrimSpace(d.Address),
			Description:   strings.TrimSpace(d.Info),
			Amenities:     trimSpaceInList(d.Amenities),
			Images:        []schema.Image{},
		}

		for category, images := range d.Images {
			for _, image := range images {
				hotel.Images = append(hotel.Images, schema.Image{URL: image.URL, Description: image.Description, Category: category})
			}
		}

		hotels = append(hotels, hotel)
	}

	return hotels, nil
}
