package bootstrap

import (
	"log"
	"github.com/spf13/viper"
)


type Env struct {
	DbUrl 					string 	`mapstructure:"DB_URL"`
	ServerAddress 			string 	`mapstructure:"SERVER_ADDRESS"`
	AccessTokenSecret 		string 	`mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret 		string 	`mapstructure:"REFRESH_TOKEN_SECRET"` 
	ContextTimeout 			int 	`mapstructure:"CONTEXT_TIMEOUT"`
	AccessTokenExpiryHour 	int 	`mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
}

func NewEnv() *Env {
	env := Env{}

	viper.SetConfigFile("../.env")

	err := viper.ReadInConfig()
	if err != nil{
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil{
		log.Fatal("enviroment can't be loaded : ", err)
	}

	return &env
}