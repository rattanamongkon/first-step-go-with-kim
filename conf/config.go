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

	envConf(`ENV_TYPE`, `dev`)

	// Database
	envConf(`DB_HOST`, `dev-datavi.cxwpcvm3mzsu.ap-southeast-1.rds.amazonaws.com:5432`)
	envConf(`DB_USERNAME`, `devdatavi`)
	envConf(`DB_PASSWORD`, `6z11KkLWZBVdGUM1Bnyh`)
	envConf(`DB_DATABASE`, `postgres`)
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
