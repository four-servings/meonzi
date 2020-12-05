package query

type (
	// Find find account
	Find struct {
		Name string
	}

	// FindResult found account result
	FindResult []struct {
		Name string `json:"name"`
	}

	findHandler interface {
		handle(query *Find) FindResult
	}

	findHandlerImplement struct {
		query AccountQuery
	}
)

func newFindHandler(query AccountQuery) findHandler {
	return &findHandlerImplement{query}
}

func (h *findHandlerImplement) handle(query *Find) FindResult {
	accounts := h.query.FindByName(query.Name)
	result := FindResult{}
	for _, account := range accounts {
		result = append(result, struct {
			Name string "json:\"name\""
		}{account.Name})
	}
	return result
}
