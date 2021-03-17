package app

import (
	"context"
	"errors"
	"github/four-servings/meonzi/account/domain"
	"github/four-servings/meonzi/pipe"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type (
	CommandBus interface {
		pipe.Bus
	}

	commandHandler struct {
		domain.AccountRepository
		domain.SocialService
	}
)

func NewCommandBus(accountRepo domain.AccountRepository, socialService domain.SocialService, timeout time.Duration) (bus CommandBus) {
	handler := commandHandler{accountRepo, socialService}
	bus = pipe.NewBusWithTimeout(timeout)

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
	id, err := ch.AccountRepository.FindNewID(ctx)
	if err != nil {
		log.WithError(err).Error("Can not get new account ID")
		return
	}

	thirdUser, err := ch.SocialService.GetUser(ctx, command.Provider, command.Token)
	if err != nil {
		log.WithError(err).Errorf("Can not fetch %s user", command.Provider)
		return
	}

	exists, _ := ch.AccountRepository.FindByProviderAndSocialID(ctx, command.Provider, command.Token)
	if exists != nil {
		err = errors.New("exists account")
		log.Error("Account is exists")
		return
	}

	account := domain.NewAccount(domain.NewAccountOptions{
		ID:           id,
		Name:         command.Name,
		AuthProvider: thirdUser.AuthProvider(),
		SocialID:     thirdUser.ID(),
	})

	account, err = ch.AccountRepository.Save(ctx, account)
	if err != nil {
		log.WithError(err).Error("Can not save account")
		return
	}

	return nil
}

type DeregisterAccountCommand struct {
	ID uuid.UUID
}

//TODO refactor, error handling on http handler
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