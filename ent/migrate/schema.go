// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// HotelsColumns holds the columns for the "hotels" table.
	HotelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "destination_id", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "latitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "longitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "address", Type: field.TypeString},
		{Name: "city", Type: field.TypeString, Nullable: true},
		{Name: "country", Type: field.TypeString, Nullable: true},
		{Name: "postal_code", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "amenities", Type: field.TypeJSON, Nullable: true},
		{Name: "images", Type: field.TypeJSON, Nullable: true},
		{Name: "booking_conditions", Type: field.TypeJSON, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// HotelsTable holds the schema information for the "hotels" table.
	HotelsTable = &schema.Table{
		Name:       "hotels",
		Columns:    HotelsColumns,
		PrimaryKey: []*schema.Column{HotelsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "hotel_id",
				Unique:  true,
				Columns: []*schema.Column{HotelsColumns[0]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		HotelsTable,
	}
)

func init() {
}