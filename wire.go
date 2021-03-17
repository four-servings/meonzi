//+build wireinject

package main

import (
	"github.com/google/wire"
	"github/four-servings/meonzi/app"
	"github/four-servings/meonzi/di"
	"time"
)

//func exampleGetAccountRepository() domain.AccountRepository {
//	wire.Build(di.ConfigSets)
//	return nil
//}

func GetApp() *app.App {
	wire.Build(
		di.ProviderSets,
		di.InfraSets,
		wire.Value(time.Second * 3),
		di.CommandBusSets,
		di.ControllerSets,
		di.RouteSets,
		app.AppSets)
	return nil
}