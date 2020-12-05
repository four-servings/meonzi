package account

import (
	"github/four-servings/meonzi/account/api"
	"github/four-servings/meonzi/account/app/command"
	"github/four-servings/meonzi/account/infra"
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
