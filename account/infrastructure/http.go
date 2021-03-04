package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github/four-servings/meonzi/account/application"
)

type handler struct {
	application.CommandBus
}

func NewRouter(commandBus application.CommandBus) {
}
