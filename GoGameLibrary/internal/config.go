package internal

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	RunEnv     string `mapstructure:"GAMELIB_RUNTIME"`
	RestPort   string `mapstructure:"GAMELIB_REST_PORT"`
	DBPort     string `mapstructure:"GAMELIB_DB_PORT"`
	DBName     string `mapstructure:"GAMELIB_DB_NAME"`
	DBUser     string `mapstructure:"GAMELIB_DB_USER"`
	DBPassword string `mapstructure:"GAMELIB_DB_PASSWORD"`
	DBHost     string `mapstructure:"GAMELIB_DB_HOST"`
}

var CFG *AppConfig

func GetAppConfig() {

	env := AppConfig{}
	viper.SetConfigFile(".env") //read the local env file
	viper.AutomaticEnv()        // override with EnvVars - if present

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't load the file .env: ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.RunEnv == "dev" {
		log.Println("The GO GAME LIBRARY app is running in a development environment")
	}

	CFG = &env // Set global, to access cfg throughout app
}
