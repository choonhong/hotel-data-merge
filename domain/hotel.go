package domain

type Hotel struct {
	ID            string
	DestinationID string
	Name          string
	Latitude      float64
	Longitude     float64
	Address       string
	City          string
	Country       string
	Details       string
	Amenities     []string
	Images        []Image
}

type Image struct {
	URL         string
	Description string
}
