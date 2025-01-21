package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type Hotel struct {
	ID                string `gorm:"primaryKey"`
	DestinationID     int
	Name              string
	Latitude          *float64
	Longitude         *float64
	Address           string
	City              string
	Country           string
	PostalCode        string
	Description       string
	Amenities         Strings `gorm:"type:VARCHAR(255)"`
	Images            Images  `gorm:"type:JSON"`
	BookingConditions Strings `gorm:"type:VARCHAR(255)"`
	UpdatedAt         time.Time
}

type Strings []string

func (o *Strings) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value cannot cast to []byte")
	}

	*o = strings.Split(string(bytes), ",")
	return nil
}

func (o Strings) Value() (driver.Value, error) {
	if len(o) == 0 {
		return nil, nil
	}

	return strings.Join(o, ","), nil
}

type Images []Image

func (i *Images) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value cannot cast to []byte")
	}

	return json.Unmarshal(bytes, i)
}

func (i Images) Value() (driver.Value, error) {
	if len(i) == 0 {
		return nil, nil
	}

	return json.Marshal(i)
}

type Image struct {
	URL         string `json:"url"`
	Description string `json:"description"`
	Category    string `json:"category"`
}
