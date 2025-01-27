// Code generated by ent, DO NOT EDIT.

package hotel

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/choonhong/hotel-data-merge/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Hotel {
	return predicate.Hotel(sql.FieldContainsFold(FieldID, id))
}

// DestinationID applies equality check predicate on the "destination_id" field. It's identical to DestinationIDEQ.
func DestinationID(v int) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldDestinationID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldName, v))
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldLatitude, v))
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldLongitude, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldAddress, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldCity, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldCountry, v))
}

// PostalCode applies equality check predicate on the "postal_code" field. It's identical to PostalCodeEQ.
func PostalCode(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldPostalCode, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldDescription, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldUpdatedAt, v))
}

// DestinationIDEQ applies the EQ predicate on the "destination_id" field.
func DestinationIDEQ(v int) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldDestinationID, v))
}

// DestinationIDNEQ applies the NEQ predicate on the "destination_id" field.
func DestinationIDNEQ(v int) predicate.Hotel {
	return predicate.Hotel(sql.FieldNEQ(FieldDestinationID, v))
}

// DestinationIDIn applies the In predicate on the "destination_id" field.
func DestinationIDIn(vs ...int) predicate.Hotel {
	return predicate.Hotel(sql.FieldIn(FieldDestinationID, vs...))
}

// DestinationIDNotIn applies the NotIn predicate on the "destination_id" field.
func DestinationIDNotIn(vs ...int) predicate.Hotel {
	return predicate.Hotel(sql.FieldNotIn(FieldDestinationID, vs...))
}

// DestinationIDGT applies the GT predicate on the "destination_id" field.
func DestinationIDGT(v int) predicate.Hotel {
	return predicate.Hotel(sql.FieldGT(FieldDestinationID, v))
}

// DestinationIDGTE applies the GTE predicate on the "destination_id" field.
func DestinationIDGTE(v int) predicate.Hotel {
	return predicate.Hotel(sql.FieldGTE(FieldDestinationID, v))
}

// DestinationIDLT applies the LT predicate on the "destination_id" field.
func DestinationIDLT(v int) predicate.Hotel {
	return predicate.Hotel(sql.FieldLT(FieldDestinationID, v))
}

// DestinationIDLTE applies the LTE predicate on the "destination_id" field.
func DestinationIDLTE(v int) predicate.Hotel {
	return predicate.Hotel(sql.FieldLTE(FieldDestinationID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldContainsFold(FieldName, v))
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldLatitude, v))
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldNEQ(FieldLatitude, v))
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldIn(FieldLatitude, vs...))
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldNotIn(FieldLatitude, vs...))
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldGT(FieldLatitude, v))
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldGTE(FieldLatitude, v))
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldLT(FieldLatitude, v))
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldLTE(FieldLatitude, v))
}

// LatitudeIsNil applies the IsNil predicate on the "latitude" field.
func LatitudeIsNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldIsNull(FieldLatitude))
}

// LatitudeNotNil applies the NotNil predicate on the "latitude" field.
func LatitudeNotNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldNotNull(FieldLatitude))
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldLongitude, v))
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldNEQ(FieldLongitude, v))
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldIn(FieldLongitude, vs...))
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldNotIn(FieldLongitude, vs...))
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldGT(FieldLongitude, v))
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldGTE(FieldLongitude, v))
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldLT(FieldLongitude, v))
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v float64) predicate.Hotel {
	return predicate.Hotel(sql.FieldLTE(FieldLongitude, v))
}

// LongitudeIsNil applies the IsNil predicate on the "longitude" field.
func LongitudeIsNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldIsNull(FieldLongitude))
}

// LongitudeNotNil applies the NotNil predicate on the "longitude" field.
func LongitudeNotNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldNotNull(FieldLongitude))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldContainsFold(FieldAddress, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldHasSuffix(FieldCity, v))
}

// CityIsNil applies the IsNil predicate on the "city" field.
func CityIsNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldIsNull(FieldCity))
}

// CityNotNil applies the NotNil predicate on the "city" field.
func CityNotNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldNotNull(FieldCity))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldContainsFold(FieldCity, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryIsNil applies the IsNil predicate on the "country" field.
func CountryIsNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldIsNull(FieldCountry))
}

// CountryNotNil applies the NotNil predicate on the "country" field.
func CountryNotNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldNotNull(FieldCountry))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldContainsFold(FieldCountry, v))
}

// PostalCodeEQ applies the EQ predicate on the "postal_code" field.
func PostalCodeEQ(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldPostalCode, v))
}

// PostalCodeNEQ applies the NEQ predicate on the "postal_code" field.
func PostalCodeNEQ(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNEQ(FieldPostalCode, v))
}

// PostalCodeIn applies the In predicate on the "postal_code" field.
func PostalCodeIn(vs ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldIn(FieldPostalCode, vs...))
}

// PostalCodeNotIn applies the NotIn predicate on the "postal_code" field.
func PostalCodeNotIn(vs ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNotIn(FieldPostalCode, vs...))
}

// PostalCodeGT applies the GT predicate on the "postal_code" field.
func PostalCodeGT(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGT(FieldPostalCode, v))
}

// PostalCodeGTE applies the GTE predicate on the "postal_code" field.
func PostalCodeGTE(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGTE(FieldPostalCode, v))
}

// PostalCodeLT applies the LT predicate on the "postal_code" field.
func PostalCodeLT(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLT(FieldPostalCode, v))
}

// PostalCodeLTE applies the LTE predicate on the "postal_code" field.
func PostalCodeLTE(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLTE(FieldPostalCode, v))
}

// PostalCodeContains applies the Contains predicate on the "postal_code" field.
func PostalCodeContains(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldContains(FieldPostalCode, v))
}

// PostalCodeHasPrefix applies the HasPrefix predicate on the "postal_code" field.
func PostalCodeHasPrefix(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldHasPrefix(FieldPostalCode, v))
}

// PostalCodeHasSuffix applies the HasSuffix predicate on the "postal_code" field.
func PostalCodeHasSuffix(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldHasSuffix(FieldPostalCode, v))
}

// PostalCodeIsNil applies the IsNil predicate on the "postal_code" field.
func PostalCodeIsNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldIsNull(FieldPostalCode))
}

// PostalCodeNotNil applies the NotNil predicate on the "postal_code" field.
func PostalCodeNotNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldNotNull(FieldPostalCode))
}

// PostalCodeEqualFold applies the EqualFold predicate on the "postal_code" field.
func PostalCodeEqualFold(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEqualFold(FieldPostalCode, v))
}

// PostalCodeContainsFold applies the ContainsFold predicate on the "postal_code" field.
func PostalCodeContainsFold(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldContainsFold(FieldPostalCode, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Hotel {
	return predicate.Hotel(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Hotel {
	return predicate.Hotel(sql.FieldContainsFold(FieldDescription, v))
}

// AmenitiesIsNil applies the IsNil predicate on the "amenities" field.
func AmenitiesIsNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldIsNull(FieldAmenities))
}

// AmenitiesNotNil applies the NotNil predicate on the "amenities" field.
func AmenitiesNotNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldNotNull(FieldAmenities))
}

// ImagesIsNil applies the IsNil predicate on the "images" field.
func ImagesIsNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldIsNull(FieldImages))
}

// ImagesNotNil applies the NotNil predicate on the "images" field.
func ImagesNotNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldNotNull(FieldImages))
}

// BookingConditionsIsNil applies the IsNil predicate on the "booking_conditions" field.
func BookingConditionsIsNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldIsNull(FieldBookingConditions))
}

// BookingConditionsNotNil applies the NotNil predicate on the "booking_conditions" field.
func BookingConditionsNotNil() predicate.Hotel {
	return predicate.Hotel(sql.FieldNotNull(FieldBookingConditions))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Hotel {
	return predicate.Hotel(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Hotel {
	return predicate.Hotel(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Hotel {
	return predicate.Hotel(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Hotel {
	return predicate.Hotel(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Hotel {
	return predicate.Hotel(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Hotel {
	return predicate.Hotel(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Hotel {
	return predicate.Hotel(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Hotel {
	return predicate.Hotel(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Hotel) predicate.Hotel {
	return predicate.Hotel(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Hotel) predicate.Hotel {
	return predicate.Hotel(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Hotel) predicate.Hotel {
	return predicate.Hotel(sql.NotPredicates(p))
}
