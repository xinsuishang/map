// Code generated by ent, DO NOT EDIT.

package tenant

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the tenant type in the database.
	Label = "tenant"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRegion holds the string denoting the region field in the database.
	FieldRegion = "region"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldAccessKey holds the string denoting the access_key field in the database.
	FieldAccessKey = "access_key"
	// FieldSecretKey holds the string denoting the secret_key field in the database.
	FieldSecretKey = "secret_key"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// Table holds the table name of the tenant in the database.
	Table = "tenants"
)

// Columns holds all SQL columns for tenant fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldRegion,
	FieldType,
	FieldAccessKey,
	FieldSecretKey,
	FieldDesc,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int32) error
)

// OrderOption defines the ordering options for the Tenant queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByRegion orders the results by the region field.
func ByRegion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegion, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByAccessKey orders the results by the access_key field.
func ByAccessKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccessKey, opts...).ToFunc()
}

// BySecretKey orders the results by the secret_key field.
func BySecretKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecretKey, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}
