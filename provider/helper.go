package provider

import (
	"encoding/json"
	"strconv"
	"strings"
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

func trimSpaceInList(list []string) []string {
	var result []string
	for _, s := range list {
		result = append(result, strings.TrimSpace(s))
	}
	return result
}
