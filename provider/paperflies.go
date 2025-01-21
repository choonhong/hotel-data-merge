package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/choonhong/hotel-data-merge/model"
)

type Paperflies struct {
	URL string
}

func (p *Paperflies) Name() string {
	return "Paperflies"
}

// see testdata/paperflies.json
type PaperfliesData struct {
	ID                string                       `json:"hotel_id"`
	DestinationID     int                          `json:"destination_id"`
	Name              string                       `json:"hotel_name"`
	Location          PaperfliesLocation           `json:"location"`
	Details           string                       `json:"details"`
	Amenities         map[string][]string          `json:"amenities"`
	Images            map[string][]PaperfliesImage `json:"images"`
	BookingConditions []string                     `json:"booking_conditions"`
}

type PaperfliesLocation struct {
	Address string `json:"address"`
	Country string `json:"country"`
}

type PaperfliesImage struct {
	Link    string `json:"link"`
	Caption string `json:"caption"`
}

func (p *Paperflies) FetchAll(ctx context.Context) ([]*model.Hotel, error) {
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

	var data []*PaperfliesData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("Decode: %w", err)
	}

	var hotels []*model.Hotel
	for _, d := range data {
		hotel := &model.Hotel{
			ID:                d.ID,
			DestinationID:     d.DestinationID,
			Name:              strings.TrimSpace(d.Name),
			Address:           strings.TrimSpace(d.Location.Address),
			Country:           strings.TrimSpace(d.Location.Country),
			Description:       strings.TrimSpace(d.Details),
			Amenities:         []string{},
			Images:            []model.Image{},
			BookingConditions: d.BookingConditions,
		}

		for _, amenities := range d.Amenities {
			hotel.Amenities = append(hotel.Amenities, trimSpaceInList(amenities)...)
		}

		for category, images := range d.Images {
			for _, image := range images {
				hotel.Images = append(hotel.Images, model.Image{URL: image.Link, Description: image.Caption, Category: category})
			}
		}

		hotels = append(hotels, hotel)
	}

	return hotels, nil
}
