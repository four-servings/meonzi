package app

import (
	"context"
	"github/four-servings/meonzi/account/domain"
	"github/four-servings/meonzi/account/infra"
	"github/four-servings/meonzi/local"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type (
	CommandBus interface {
		local.Bus
	}

	commandHandler struct {
		domain.AccountRepository
		KakaoAdapter
		GoogleAdapter
	}
)

func NewCommandBus(accountRepo domain.AccountRepository, kakaoAdapter KakaoAdapter, googleAdapter GoogleAdapter, timeout time.Duration) (bus CommandBus) {
	handler := commandHandler{accountRepo, kakaoAdapter, googleAdapter}
	bus = CommandBus(local.NewBusWithTimeout(timeout))

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return bus.RegistryHandler(RegisterAccountCommand{}, handler.RegisterAccountHandle)
	})
	g.Go(func() error {
		return bus.RegistryHandler(DeregisterAccountCommand{}, handler.DeregisterAccountHandle)
	})
	err := g.Wait()
	if err != nil {
		panic(err)
	}

	return
}

type RegisterAccountCommand struct {
	Name     string
	Token    string
	Provider domain.AuthProvider
}

func (ch *commandHandler) RegisterAccountHandle(ctx context.Context, command RegisterAccountCommand) (err error) {
	id, err := ch.AccountRepository.FindNewID()
	if err != nil {
		log.WithError(err).Error("Can not get new account ID")
		return
	}

	var thirdUser ThirdUser

	if command.Provider == domain.KakaoServiceProviderKey {
		thirdUser, err = ch.KakaoAdapter.GetUser(ctx, command.Token)
		if err != nil {
			log.WithError(err).Error("Can not fetch kakao user")
			return
		}
	}

	if command.Provider == domain.GoogleServiceProviderKey {
		thirdUser, err = ch.GoogleAdapter.GetUser(ctx, command.Token)
		if err != nil {
			log.WithError(err).Error("Can not fetch google user")
		}
	}

	account := domain.NewAccount(domain.NewAccountOptions{
		ID:           id,
		Name:         command.Name,
		AuthProvider: command.Provider,
		SocialID:     thirdUser.ID(),
	})

	account, err = ch.AccountRepository.Save(ctx, account)
	if err != nil {
		log.WithError(err).Error("Can not save account")
	}

	return nil
}

func (ch *commandHandler) getSocialAdpater(provider domain.AuthProvider, token string) SocialAdapter {
	switch provider {
	case domain.GoogleServiceProviderKey:
		return ch.GoogleAdapter
	case domain.KakaoServiceProviderKey:
		return ch.KakaoAdapter
	}

	return infra.UnknownAdapter
}

type DeregisterAccountCommand struct {
	ID uuid.UUID
}

func (ch *commandHandler) DeregisterAccountHandle(ctx context.Context, command DeregisterAccountCommand) error {
	account, err := ch.AccountRepository.FindByID(ctx, command.ID)
	if err != nil {
		log.WithError(err).Error("Can not found account")
	}

	account.Deregister()

	account, err = ch.AccountRepository.Save(ctx, account)
	if err != nil {
		log.WithError(err).Error("Can not save account")
	}

	return nil
}
