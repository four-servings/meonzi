package di

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github/four-servings/meonzi/account/infrastructure"
	"github/four-servings/meonzi/config"
	"github/four-servings/meonzi/ent"
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

// config
var ConfigSets = wire.NewSet(
	config.GetDBConn,
	)

// providers
var ProviderSets = wire.NewSet(
	ProviderDatabase,
	ProviderAccountTable,
	)

// repositories
var RepositorySets = wire.NewSet(
	infrastructure.NewAccountRepository,
	)
