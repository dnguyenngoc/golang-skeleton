package settings

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/apot-group/golang-skeleton/x-api/src/loggers"
)

type (
	Config struct {
		Rabbitmq rabbitmq
		Storage  storage
	}

	rabbitmq struct {
		Host  string
		Post  string
		User  string
		Pass  string
		Vhost string
	}
	storage struct {
		Image string
	}
)

// config is a global variable to hold config
var ConfEnv *Config

func InitConfig() {
	loggers.InfoLogger.Println("Load toml file config ...")
	f := "./environment.toml"
	if _, err := os.Stat(f); err != nil {
		loggers.WarningLogger.Println("wrong env file path -> try to other path.")
		f = "./settings/environment.toml"
	}
	_, err := toml.DecodeFile(f, &ConfEnv)
	if err != nil {
		loggers.ErrorLogger.Fatal(err)
		os.Exit(1)
		log.Panic(err)
	}
	loggers.InfoLogger.Println("Init config is Decoded!")
}
