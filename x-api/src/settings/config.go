package settings

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
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

	fmt.Println("==========================")
	fmt.Println("Load toml file config ...")
	fmt.Println("==========================")
	f := "./environment.toml"
	if _, err := os.Stat(f); err != nil {
		f = "./settings/environment.toml"
	}
	_, err := toml.DecodeFile(f, &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Decoded!")

}
