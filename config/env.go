package config

import (
	"fmt"
	"github/four-servings/meonzi/di"
	"net/url"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/google/wire"
)

var config struct {
	DbUser string `env:"DATABASE_USER" envDefault:"root"`
	DbPass string `env:"DATABASE_PASS" envDefault:"test"`
	DbHost string `env:"DATABASE_HOST" envDefault:"localhost"`
	DbName string `env:"DATABASE_NAME" envDefault:"meonzi"`
	DbPort uint16 `env:"DATABASE_PORT" envDefault:"3306"`

	//RedisAddr string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	//RedisPass string `env:"REDIS_PASS"`
	//RedisUser string `env:"REDIS_USER" envDefault:"default"`
	//RedisDB int `env:"REDIS_DB" envDefault:"10"`

	Port         uint16 `env:"PORT" envDefault:"5000"`
	IsProduction bool   `env:"PRODUCTION" envDefault:"true"`
	TimeZone     string `env:"TZ" envDefault:"UTC"`
	IsDebug      bool   `env:"DEBUG" envDefault:"true"`
}

func init() {
	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}
}

func GetDBConn() di.DBConn {
	val := url.Values{}
	val.Add("parseTime", "true")
	val.Add("loc", time.UTC.String())
	return di.DBConn(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.DbUser,
		config.DbPass,
		config.DbHost,
		config.DbPort,
		config.DbName,
		val.Encode()))
}

// config
var ConfigSets = wire.NewSet(
	GetDBConn,
)
