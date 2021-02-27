package application

import "github/four-servings/meonzi/account/infrastructure"

type (
	FindByID struct {
		ID string
	}
	FindByIDHandler struct {
		query *infrastructure.Query
	}
)


func (h *FindByIDHandler) handle(query FindByID) interface{} {
	retrun h.query.FindByID(query.ID)
}

