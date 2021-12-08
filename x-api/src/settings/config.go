package settings

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/apot-group/golang-skeleton/x-api/src/loggers"
)

type (
	Config struct {
		Rabbitmq map[string]rabbitmq
	}

	rabbitmq struct {
		Host  string
		Post  string
		User  string
		Pass  string
		Vhost string
	}
)

// config is a global variable to hold config
var config *Config

func InitConfig() {

	loggers.InfoLogger.Println("Load toml file config ...")
	f := "./environment.toml"
	if _, err := os.Stat(f); err != nil {
		loggers.WarningLogger.Println("wrong env file path -> try to other path.")
		f = "./settings/environment.toml"
	}
	_, err := toml.DecodeFile(f, &config)
	if err != nil {
		// fmt.Fprintln(os.Stderr, err)
		loggers.ErrorLogger.Fatal(err)
		os.Exit(1)
	}
	loggers.InfoLogger.Println("Init config is Decoded!")

}
