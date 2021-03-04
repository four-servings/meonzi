package infrastructure

import (
	"github/four-servings/meonzi/ent"
	"github/four-servings/meonzi/ent/schema"
)

type (
	GetAccountBySocial struct {
		SocialType schema.SocialType
		SocialId   string
	}

	Query interface {

	}

	queryImpl struct {
		cli *ent.AccountClient
	}
)

func (q *queryImpl) FindByID(id string) {
	// SELECT t....
}
