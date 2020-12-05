package command

import (
	"github/four-servings/meonzi/accounts/domain"
)

type (
	// CreateAccount create account command
	CreateAccount struct {
		Name string
	}

	createAccountHandler interface {
		handle(command *CreateAccount)
	}

	createAccountHandlerImplement struct {
		repository domain.AccountRepository
	}
)

func newCreateAccountHandler(repository domain.AccountRepository) createAccountHandler {
	return &createAccountHandlerImplement{repository}
}

func (h *createAccountHandlerImplement) handle(command *CreateAccount) {
	id := h.repository.FindNewID()
	account := domain.NewAccount(id, command.Name)
	h.repository.Save(account)
}
