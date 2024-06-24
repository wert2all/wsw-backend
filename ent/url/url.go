// Code generated by ent, DO NOT EDIT.

package enturl

import (
	"fmt"
	"wsw/backend/domain/url"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the url type in the database.
	Label = "url"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldAPIURLID holds the string denoting the api_url_id field in the database.
	FieldAPIURLID = "api_url_id"
	// Table holds the table name of the url in the database.
	Table = "urls"
)

// Columns holds all SQL columns for url fields.
var Columns = []string{
	FieldID,
	FieldURL,
	FieldStatus,
	FieldAPIURLID,
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

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s url.Status) error {
	switch s {
	case "success", "error", "pending":
		return nil
	default:
		return fmt.Errorf("enturl: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Url queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByAPIURLID orders the results by the api_url_id field.
func ByAPIURLID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAPIURLID, opts...).ToFunc()
}
