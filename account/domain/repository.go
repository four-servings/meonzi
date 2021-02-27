package domain

import (
	"context"
	"github.com/google/uuid"
	"github/four-servings/meonzi/ent"
	"github/four-servings/meonzi/ent/schema"
)

type AccountRepository interface {
	Create(ctx context.Context, data *ent.Account) (err error)
	GetBySocial(ctx context.Context, typ schema.SocialType, id string) (res *ent.Account, err error)
	Update(ctx context.Context, data *ent.Account) (err error)
	Delete(ctx context.Context, id uuid.UUID) (err error)
}
