//+build wireinject

package main

import (
	"github.com/google/wire"
	"github/four-servings/meonzi/account/domain"
	"github/four-servings/meonzi/account/infrastructure"
	"github/four-servings/meonzi/di"
)


func exampleGetAccountRepository() domain.AccountRepository {
	wire.Build(
		di.ConfigSets,
		di.ProviderSets,
		infrastructure.NewAccountRepository,
	)
	return nil
}
