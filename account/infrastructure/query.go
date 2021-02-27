package infrastructure

import (
	"context"
	"github/four-servings/meonzi/ent/schema"
)

type (
	GetAccountBySocial struct {
		Ctx        context.Context
		SocialType schema.SocialType
		SocialId   string
	}

	Query struct{}
)

func (q *Query) FindByID(id string) {
	// SELECT t....
}
