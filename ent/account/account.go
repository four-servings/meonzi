// Code generated by entc, DO NOT EDIT.

package account

import (
	"time"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSocialType holds the string denoting the social_type field in the database.
	FieldSocialType = "social_type"
	// FieldSocialID holds the string denoting the social_id field in the database.
	FieldSocialID = "social_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLastAccessedAt holds the string denoting the last_accessed_at field in the database.
	FieldLastAccessedAt = "last_accessed_at"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"

	// Table holds the table name of the account in the database.
	Table = "accounts"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldSocialType,
	FieldSocialID,
	FieldName,
	FieldLastAccessedAt,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
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
	// SocialIDValidator is a validator for the "social_id" field. It is called by the builders before save.
	SocialIDValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() time.Time
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() time.Time
)