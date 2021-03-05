package infrastructure

import (
	"github/four-servings/meonzi/account/interfaces"
	"net/http"
)

type handler struct {
	interfaces.Controller
}

func NewRouter(e *echo.Echo, controller interfaces.Controller) {
	handler := handler{controller}
	e := echo.New()
	e.POST("/account", handler.registerAccount)
}

func (h *handler) registerAccount(ctx echo.Context) error {
	var binder struct {
		Token    string `json:"token"`
		Provider string `json:"provider"`
		Name     string `json:"name"`
	}
	err := ctx.Bind(&binder)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	ctx.Validate(binder)
	dto := interfaces.RegisterAccountDTO{
		Token:    binder.Token,
		Name:     binder.Name,
		Provider: binder.Provider,
	}
	h.Controller.RegisterAccount(dto)
	return nil
}
