package infra

import (
	"github.com/labstack/echo/v4"
	"github/four-servings/meonzi/account/interfaces"
	"net/http"
)

type handler struct {
	interfaces.Controller
}

type Routing func()

func NewRoute(e *echo.Echo, controller interfaces.Controller) Routing {
	return func() {
		handler := handler{controller}
		e.POST("/account", handler.registerAccount)
		e.POST("/auth", )
	}
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

	dto := interfaces.RegisterAccountDTO{
		Token:    binder.Token,
		Name:     binder.Name,
		Provider: binder.Provider,
	}
	//TODO error handling
	h.Controller.RegisterAccount(dto)
	return nil
}

func (h *handler) getAccessToken(ctx echo.Context) error {
	var binder struct {
		Token string `json:"token"`
		Provider string `json:"provider"`
	}

	err := ctx.Bind(&binder)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	token, err := h.Controller.SignIn(interfaces.SignInAccountDTO{
		Token:    binder.Token,
		Provider: binder.Provider,
	})
	if err != nil {
		//TODO error handling
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, token)
}