package infrastructure

import (
	"context"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github/four-servings/meonzi/account/domain"
	"github/four-servings/meonzi/ent"
	"github/four-servings/meonzi/ent/account"
	"github/four-servings/meonzi/ent/schema"
	"github/four-servings/meonzi/util/arr"
	"time"
)

type repo struct {
	cli *ent.AccountClient
}

func NewAccountRepository(cli *ent.AccountClient) domain.AccountRepository {
	return &repo{cli}
}

func (r *repo) Create(ctx context.Context, data *ent.Account) (err error) {
	res, err := r.cli.Create().
		SetID(data.ID).
		SetSocialType(data.SocialType).
		SetSocialID(data.SocialID).
		SetName(data.Name).
		SetLastAccessedAt(data.LastAccessedAt).
		Save(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"data": data,
		}).WithError(err).Error("account create exception")
	} else {
		*data = *res
	}
	return
}

func (r *repo) GetBySocial(ctx context.Context, typ schema.SocialType, id string) (res *ent.Account, err error) {
	res, err = r.cli.Query().
		Where(
			account.And(
				account.SocialType(typ),
				account.SocialID(id),
			),
		).First(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"args": arr.Flatten(typ, id),
		}).WithError(err).Error("account get by social exception")
	}

	return
}


func (r *repo) Update(ctx context.Context, data *ent.Account) (err error) {
	res, err := r.cli.UpdateOneID(data.ID).
		SetName(data.Name).
		Save(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"data": data,
		}).WithError(err).Error("account update exception")
	} else {
		*data = *res
	}

	return
}

func (r *repo) Delete(ctx context.Context, id uuid.UUID) (err error) {
	err = r.cli.UpdateOneID(id).
		SetDeleteAt(time.Now()).
		Exec(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"id": id,
		}).WithError(err).Error("account delete exception")
	}
	return
}