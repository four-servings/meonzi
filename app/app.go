package app

import (
	"fmt"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github/four-servings/meonzi/account/infra"
	"github/four-servings/meonzi/ent"
)

var AppSets = wire.NewSet(
	GetDBConn,
	Routing,
	NewApp,
	)

type setRoute func()

type App struct {
	db *ent.Client
	e *echo.Echo
}

func Routing(r1 infra.Routing) setRoute {
	return func() {
		r1()
	}
}

func NewApp(db *ent.Client, e *echo.Echo, setRoute setRoute) (app *App) {
	app = &App{db, e}
	setRoute()
	return
}

func (a *App) Start() {
	defer a.close()
	e := a.e
	e.Start(fmt.Sprintf(":%d", config.Port))
}

func (a *App) close() {
	err := a.db.Close()
	if err != nil {
		log.WithError(err).Error("database close exception")
	}
}