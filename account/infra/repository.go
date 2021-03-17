package infra

import (
	"context"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github/four-servings/meonzi/account/domain"
	"github/four-servings/meonzi/ent"
	"github/four-servings/meonzi/ent/account"
	"github/four-servings/meonzi/ent/schema"
	"github/four-servings/meonzi/util/arr"
)

type repo struct {
	cli *ent.AccountClient
}


func NewAccountRepository(cli *ent.AccountClient) domain.AccountRepository {
	return &repo{cli}
}

func (r *repo) FindNewID(ctx context.Context) (newId uuid.UUID, err error) {
	id := uuid.New()
	exists, _ := r.cli.Get(ctx, id)
	if exists != nil {
		log.Error("account find new id exception")
		err = errors.New("account find new id exception")
		return
	}
	newId = id
	return
}

func (r *repo) Save(ctx context.Context, account domain.Account) (saved domain.Account, err error) {
	data := account.Data()
	entity, _ := r.cli.Get(ctx, data.ID)
	var save func(context.Context) (*ent.Account, error)
	if entity != nil {
		save = entity.Update().
			SetName(data.Name).
			SetNillableDeleteAt(data.DeletedAt).Save
	} else {
		save = r.cli.Create().
			SetID(data.ID).
			SetSocialType(convertAuthProviderToSocialType(data.AuthProvider)).
			SetSocialID(data.SocialID).
			SetName(data.Name).Save
	}

	entity, err = save(ctx)
	if err != nil {
		log.WithError(err).Error("account save exception")
		return
	}
	saved = convertEntityToDomain(entity)
	return
}

func (r *repo) FindByID(ctx context.Context, id uuid.UUID) (res domain.Account, err error) {
	acc, err := r.cli.Get(ctx, id)
	if err != nil {
		log.WithError(err).Error("account find by id exception")
		return
	}

	res = convertEntityToDomain(acc)
	return
}

func (r *repo) FindByProviderAndSocialID(ctx context.Context, provider domain.AuthProvider, socialID string) (res domain.Account, err error) {
	acc, err := r.cli.Query().
		Where(account.And(
			account.SocialType(convertAuthProviderToSocialType(provider)),
			account.SocialID(socialID),
			)).First(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"args": arr.Flatten(provider, socialID),
		}).WithError(err).Error("account get by social exception")
		return
	}

	res = convertEntityToDomain(acc)
	return
}

func convertAuthProviderToSocialType(provider domain.AuthProvider) schema.SocialType {
	switch provider {
	case domain.KakaoServiceProviderKey:
		return schema.SocialTypeKakao
	case domain.GoogleServiceProviderKey:
		return schema.SocialTypeGoogle
	default:
		return schema.SocialTypeUnknown
	}
}

func convertEntityToDomain(from *ent.Account) domain.Account {
	return domain.ReconstituteAccount(domain.ReconstituteAccountOptions{
		ID:           from.ID,
		Name:         from.Name,
		AuthProvider: convertSocialTypeToAuthProvider(from.SocialType),
		SocialID:     from.SocialID,
		CreatedAt:    from.CreateAt,
		UpdatedAt:    from.UpdateAt,
		DeletedAt:    from.DeleteAt,
	})
}

func convertSocialTypeToAuthProvider(socialType schema.SocialType) domain.AuthProvider {
	switch socialType {
	case schema.SocialTypeKakao:
		return domain.KakaoServiceProviderKey
	case schema.SocialTypeGoogle:
		return domain.GoogleServiceProviderKey
	default:
		return "-"
	}
}