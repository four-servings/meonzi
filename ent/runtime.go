// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github/four-servings/meonzi/ent/account"
	"github/four-servings/meonzi/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescSocialID is the schema descriptor for social_id field.
	accountDescSocialID := accountFields[2].Descriptor()
	// account.SocialIDValidator is a validator for the "social_id" field. It is called by the builders before save.
	account.SocialIDValidator = accountDescSocialID.Validators[0].(func(string) error)
	// accountDescName is the schema descriptor for name field.
	accountDescName := accountFields[3].Descriptor()
	// account.NameValidator is a validator for the "name" field. It is called by the builders before save.
	account.NameValidator = accountDescName.Validators[0].(func(string) error)
	// accountDescCreateAt is the schema descriptor for create_at field.
	accountDescCreateAt := accountFields[5].Descriptor()
	// account.DefaultCreateAt holds the default value on creation for the create_at field.
	account.DefaultCreateAt = accountDescCreateAt.Default.(func() time.Time)
	// accountDescUpdateAt is the schema descriptor for update_at field.
	accountDescUpdateAt := accountFields[6].Descriptor()
	// account.DefaultUpdateAt holds the default value on creation for the update_at field.
	account.DefaultUpdateAt = accountDescUpdateAt.Default.(func() time.Time)
	// account.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	account.UpdateDefaultUpdateAt = accountDescUpdateAt.UpdateDefault.(func() time.Time)
}