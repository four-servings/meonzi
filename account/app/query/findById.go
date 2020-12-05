package query

type (
	// FindByID find account by id
	FindByID struct {
		ID string
	}

	// FindByIDResult query result for findByID
	FindByIDResult struct {
		ID   string
		Name string
	}

	findByIDHandler interface {
		handle(query *FindByID) FindByIDResult
	}

	findByIDHandlerImplement struct {
		query AccountQuery
	}
)

func newFindByIDHandler(query AccountQuery) findByIDHandler {
	return &findByIDHandlerImplement{query}
}

func (h *findByIDHandlerImplement) handle(query *FindByID) FindByIDResult {
	account := h.query.FindByID(query.ID)
	return struct {
		ID   string
		Name string
	}{account.ID, account.Name}
}
