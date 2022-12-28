package connection

import (
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

var db *pg.DB

func GetConnectionDB() *pg.DB {
	if db == nil {
		db = pg.Connect(&pg.Options{
			Addr:     viper.GetString(`DB_HOST`),
			User:     viper.GetString(`DB_USERNAME`),
			Password: viper.GetString(`DB_PASSWORD`),
			Database: viper.GetString(`DB_DATABASE`),
		})
	}
	return db
}
