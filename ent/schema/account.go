package schema

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

type SocialType uint8

const (
	SocialTypeUnknown SocialType = iota
	SocialTypeKakao
	SocialTypeGoogle
)

var (
	socialTypeToString = []string{
		"unknown",
		"kakao",
		"google",
	}

	socialTypeFromString = map[string]SocialType{
		"kakao":  SocialTypeKakao,
		"google": SocialTypeGoogle,
	}
)

func (s SocialType) String() string {
	if len(socialTypeToString) <= int(s) {
		return "unknown"
	}

	return socialTypeToString[int(s)]
}

func (s SocialType) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.String() + "\""), nil
}

func (s *SocialType) UnmarshalJSON(data []byte) error {
	v, ok := socialTypeFromString[strings.Trim(string(data), "\"")]
	if !ok {
		*s = SocialTypeUnknown
	} else {
		*s = v
	}

	return nil
}

// domain model, orm => database entity

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable(),
		field.Uint8("social_type").GoType(SocialTypeUnknown).Immutable(),
		field.String("social_id").MaxLen(30).Immutable(),
		field.String("name").MaxLen(30),
		field.Time("create_at").Default(time.Now).Immutable(),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("delete_at").Optional().Nillable(),
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("social_type", "social_id").Unique(),
		index.Fields("delete_at"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "account"},
	}
}
