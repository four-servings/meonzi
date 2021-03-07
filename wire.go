//+build wireinject

package main

import (
	"github/four-servings/meonzi/account/domain"
	"github/four-servings/meonzi/account/infra"
	"github/four-servings/meonzi/config"
	"github/four-servings/meonzi/di"

	"github.com/google/wire"
)

func exampleGetAccountRepository() domain.AccountRepository {
	wire.Build(
		config.ConfigSets,
		di.ProviderSets,
		infra.NewAccountRepository,
	)
	return nil
}
