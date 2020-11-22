package accounts

import (
	"github/four-servings/meonzi/accounts/api"
	"github/four-servings/meonzi/accounts/app/command"
	"github/four-servings/meonzi/accounts/infra"
	"github/four-servings/meonzi/setup"
	"net/http"
)

func init() {
	dbConnection := setup.GetDatabaseConnection()
	repository := infra.NewRepository(dbConnection)
	commandBus := command.NewBus(repository)
	controller := api.NewController(commandBus)

	http.HandleFunc("/accounts", controller.Handle)
}
