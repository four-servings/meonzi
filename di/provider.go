package di

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	accountApp "github/four-servings/meonzi/account/app"
	"github/four-servings/meonzi/account/infra"
	"github/four-servings/meonzi/account/interfaces"
	"github/four-servings/meonzi/ent"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

func ProviderDatabase(conn DBConn) *ent.Client {
	cli, err := ent.Open("mysql", string(conn))
	if err != nil {
		panic(err)
	}
	if err = cli.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	return cli
}

func ProviderAccountTable(cli *ent.Client) *ent.AccountClient {
	return cli.Account
}

func ProviderEcho() (e *echo.Echo) {
	e = echo.New()
	return
}

func ProviderValidator() (v *validator.Validate) {
	v = validator.New()
	return
}

var ProviderSets = wire.NewSet(
	ProviderDatabase,
	ProviderAccountTable,
	ProviderEcho,
	ProviderValidator,
)

var InfraSets = wire.NewSet(
	infra.NewAccountRepository,
	infra.NewKakaoAdapter,
	infra.NewGoogleAdapter,
	infra.NewSocialService,
	infra.NewAuthService,
)

var CommandBusSets = wire.NewSet(
	accountApp.NewCommandBus,
)

var ControllerSets = wire.NewSet(
	interfaces.NewAccountController,
)

var RouteSets = wire.NewSet(
	infra.NewRoute,
)