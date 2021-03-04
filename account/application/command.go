package application

import (
	"context"
	"github/four-servings/meonzi/account/domain"
	"github/four-servings/meonzi/account/social/google"
	"github/four-servings/meonzi/account/social/kakao"
	"github/four-servings/meonzi/local"
	"golang.org/x/sync/errgroup"
	"time"
)

type (

	CommandBus interface {
		local.Bus
	}

	commandHandler struct {
		kakao.User
		google.Auth
		domain.AccountRepository
	}
)

func NewCommandBus(kUser kakao.User, gAuth google.Auth, accountRepo domain.AccountRepository, eventBus, timeout time.Duration) (bus CommandBus) {
	handler := commandHandler{kUser, gAuth, accountRepo}
	bus = CommandBus(local.NewBusWithTimeout(timeout))

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return bus.RegistryHandler(domain.RegisterAccountCommand{}, handler.RegisterAccountHandle)
	})
	g.Go(func() error {
		return bus.RegistryHandler(domain.UpdateAccountCommand{}, handler.UpdateAccountHandle)
	})
	g.Go(func() error {
		return bus.RegistryHandler(domain.RemoveAccountCommand{}, handler.RemoveAccountHandle)
	})
	err := g.Wait()
	if err != nil {
		panic(err)
	}

	return
}

func (ch *commandHandler) RegisterAccountHandle(ctx context.Context, command domain.RegisterAccountCommand) error {
	return nil
}

func (ch *commandHandler) UpdateAccountHandle(ctx context.Context, command domain.UpdateAccountCommand) error {
	return nil
}

func (ch *commandHandler) RemoveAccountHandle(ctx context.Context, command domain.RemoveAccountCommand) error {
	return nil
}
