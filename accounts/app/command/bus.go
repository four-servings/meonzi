package command

import (
	"errors"
	"github/four-servings/meonzi/accounts/infra"
)

type (
	// Bus command bus interface
	Bus interface {
		Handle(command interface{})
	}

	// bustImplement command bus implement
	bustImplement struct {
		createAccountHandler createAccountHandler
	}
)

// NewBus create command bus instance
func NewBus(repository infra.AccountRepository) Bus {
	createAccountHandler := newCreateAccountHandler(repository)
	return &bustImplement{createAccountHandler}
}

// Handle handle given command
func (b *bustImplement) Handle(givenCommand interface{}) {
	switch givenCommand := givenCommand.(type) {
	case *CreateAccount:
		b.createAccountHandler.handle(givenCommand)
	default:
		panic(errors.New("invalid command"))
	}
}
