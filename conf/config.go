package conf

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Init() {
	godotenv.Load("app.env")
	loadENV()
}

func loadENV() {

	envConf(`ENV_TYPE`, `local`)

	// Database
	envConf(`DB_HOST`, `127.0.0.1:5432`)
	envConf(`DB_USERNAME`, `postgres`)
	envConf(`DB_PASSWORD`, `postgres`)
	envConf(`DB_DATABASE`, `myTest-datavi`)
	envConf(`CLIENT_ORIGIN`, `http://localhost:3000`)
	envConf(`TOKEN_EXPIRED_IN`, `60m`)
	envConf(`TOKEN_MAXAGE`, `60`)
	envConf(`TOKEN_SECRET`, `mykey`)

}

func envConf(key string, initValue any) {
	val, ok := os.LookupEnv(key)
	if !ok {
		viper.SetDefault(key, initValue)
	} else {
		viper.SetDefault(key, val)
	}
}
