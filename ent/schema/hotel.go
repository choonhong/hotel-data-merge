package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Hotel holds the schema definition for the Hotel entity.
type Hotel struct {
	ent.Schema
}

// Fields of the Hotel.
func (Hotel) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Int("destination_id"),
		field.String("name"),
		field.Float("latitude").Nillable().Optional(),
		field.Float("longitude").Nillable().Optional(),
		field.String("address"),
		field.String("city").Optional(),
		field.String("country").Optional(),
		field.String("postal_code").Optional(),
		field.String("description").Optional(),
		field.Strings("amenities").Optional(),
		field.JSON("images", []Image{}).Optional(),
		field.Strings("booking_conditions").Optional(),
		field.Time("updated_at").Default(time.Now),
	}
}

// Indexes of the Hotel.
func (Hotel) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

type Image struct {
	URL         string `json:"url"`
	Description string `json:"description"`
	Category    string `json:"category"`
}
