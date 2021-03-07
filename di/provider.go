package di

import (
	"context"
	"github/four-servings/meonzi/account/infra"
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

// providers
var ProviderSets = wire.NewSet(
	ProviderDatabase,
	ProviderAccountTable,
)

// repositories
var RepositorySets = wire.NewSet(
	infra.NewAccountRepository,
)
