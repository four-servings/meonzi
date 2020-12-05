package account

import (
	"github/four-servings/meonzi/account/api"
	"github/four-servings/meonzi/account/app/command"
	"github/four-servings/meonzi/account/app/query"
	"github/four-servings/meonzi/account/infra"
	"github/four-servings/meonzi/setup"
	"net/http"
)

func init() {
	dbConnection := setup.GetDatabaseConnection()
	commandBus := command.NewBus(infra.NewRepository(dbConnection))
	queryBus := query.NewBus(infra.NewQuery(dbConnection))
	controller := api.NewController(commandBus, queryBus)

	http.HandleFunc("/accounts", controller.Handle)
}
