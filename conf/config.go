package conf

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Init() {
	godotenv.Load(".env")
	loadENV()
}

func loadENV() {

	envConf(`ENV_TYPE`, `local`)

	// Database
	envConf(`DB_HOST`, `127.0.0.1:5432`)
	envConf(`DB_USERNAME`, `chaiwat`)
	envConf(`DB_PASSWORD`, ``)
	envConf(`DB_DATABASE`, `postgres`)
}

func envConf(key string, initValue any) {
	val, ok := os.LookupEnv(key)
	if !ok {
		viper.SetDefault(key, initValue)
	} else {
		viper.SetDefault(key, val)
	}
}
