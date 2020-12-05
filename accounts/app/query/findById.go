package query

import (
	"errors"
	"github/four-servings/meonzi/accounts/domain"
	"time"
)

type (
	// FindByID find account by id
	FindByID struct {
		ID string
	}

	// FindByIDResult query result for findByID
	FindByIDResult struct {
		ID             string
		Name           string
		LastAccessedAt time.Time
		CreatedAt      time.Time
		UpdatedAt      time.Time
		DeletedAt      *time.Time
	}

	findByIDHandler interface {
		handle(query *FindByID) FindByIDResult
	}

	findByIDHandlerImplement struct {
		repository domain.AccountRepository
	}
)

func newFindByIDHandler(repository domain.AccountRepository) findByIDHandler {
	return &findByIDHandlerImplement{repository}
}

func (h *findByIDHandlerImplement) handle(query *FindByID) FindByIDResult {
	account := h.repository.FindByID(query.ID)
	if account.ID() == "" {
		panic(errors.New("not found"))
	}
	return FindByIDResult{account.ID(), account.Name(), account.LastAccessedAt(), account.CreatedAt(), account.UpdatedAt(), account.DeletedAt()}
}
