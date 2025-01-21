package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/choonhong/hotel-data-merge/model"
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

func (p *Patagonia) FetchAll(ctx context.Context) ([]*model.Hotel, error) {
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

	var data []*PatagoniaData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("Decode: %w", err)
	}

	var hotels []*model.Hotel
	for _, d := range data {
		hotel := &model.Hotel{
			ID:            d.ID,
			DestinationID: d.DestinationID,
			Name:          strings.TrimSpace(d.Name),
			Latitude:      d.Latitude.Value,
			Longitude:     d.Longitude.Value,
			Address:       strings.TrimSpace(d.Address),
			City:          "",
			Country:       "",
			PostalCode:    "",
			Description:   strings.TrimSpace(d.Info),
			Amenities:     trimSpaceInList(d.Amenities),
			Images:        []model.Image{},
		}

		for category, images := range d.Images {
			for _, image := range images {
				hotel.Images = append(hotel.Images, model.Image{URL: image.URL, Description: image.Description, Category: category})
			}
		}

		hotels = append(hotels, hotel)
	}

	return hotels, nil
}
