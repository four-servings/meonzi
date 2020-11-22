package command

import (
	"github/four-servings/meonzi/accounts/domain"
	"github/four-servings/meonzi/accounts/infra"
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
		repository infra.AccountRepository
	}
)

func newCreateAccountHandler(repository infra.AccountRepository) createAccountHandler {
	return &createAccountHandlerImplement{repository}
}

func (h *createAccountHandlerImplement) handle(command *CreateAccount) {
	id := h.repository.FindNewID()
	account := domain.NewAccount(id, command.Name)
	h.repository.Save(account)
}
