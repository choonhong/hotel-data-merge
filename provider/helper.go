package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Custom type for Latitude and Longitude
type Float64OrString struct {
	Value *float64
}

func (f *Float64OrString) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch value := v.(type) {
	case float64:
		f.Value = &value
	case string:
		if value == "" {
			f.Value = nil
		} else {
			parsedValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil
			}
			f.Value = &parsedValue
		}
	}

	return nil
}

func FetchDataFromURL(ctx context.Context, url string) (io.ReadCloser, error) {
	// Call Acme API
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequestWithContext: %w", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Do: %w", err)
	}
	return resp.Body, nil
}
