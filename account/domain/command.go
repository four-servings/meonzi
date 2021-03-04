package domain

import (
	"github.com/google/uuid"
	"github/four-servings/meonzi/ent/schema"
)

type (
	RegisterAccountCommand struct {
		SocialType schema.SocialType
		AccessToken string
		Name string
	}

	UpdateAccountCommand struct {
		Id uuid.UUID
		Name string
	}

	RemoveAccountCommand struct {
		Id uuid.UUID
	}
)