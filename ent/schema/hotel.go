package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/choonhong/hotel-data-merge/domain"
)

type Hotel struct {
	ent.Schema
}

// Fields of the Hotel.
func (Hotel) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("destination_id"),
		field.String("name"),
		field.Float("latitude"),
		field.Float("longitude"),
		field.String("address"),
		field.String("city"),
		field.String("country"),
		field.String("details"),
		field.Strings("amenities"),
		field.JSON("images", []domain.Image{}),
		field.Time("updated_at").Default(time.Now()),
	}
}

func (Hotel) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("destination_id"),
	}
}
