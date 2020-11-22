package setup

import (
	"github/four-servings/dropit-backend/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetDatabaseConnection get database connection
func GetDatabaseConnection() *gorm.DB {
	dbConfig := config.Database{}

	user := dbConfig.User()
	password := dbConfig.Password()
	host := dbConfig.Host()
	port := dbConfig.Port()
	name := dbConfig.Name()

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?parseTime=true"

	connection, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      logger.Silent,
				},
			),
		},
	)
	if err != nil {
		panic(err)
	}

	return connection
}
