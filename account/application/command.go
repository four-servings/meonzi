package application

import (
	"context"
	"github.com/google/uuid"
	"github/four-servings/meonzi/ent/schema"
)

type (
	RegisterAccountCommand struct {
		Ctx context.Context
		Id uuid.UUID
		SocialType schema.SocialType
		AccessToken string
	}

	UpdateAccountCommand struct {
		Ctx context.Context
		Id uuid.UUID
	}

	RemoveAccountCommand struct {
		Ctx context.Context
		Id uuid.UUID
	}

)
