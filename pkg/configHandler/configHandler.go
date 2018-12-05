package configHandler

import (
	"flag"
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port              int
	Prefix            string
	Token             string
	Connection_String string
}

var (
	Config     Configuration
	enviroment string
	path       string
)

func init() {
	flag.StringVar(&enviroment, "env", "", "Enviroment")
	//flag.StringVar(&Password, "p", "", "Account Password")
	flag.Parse()

	getPath := func(i interface{}) string {
		switch enviroment {
		case "prod":
			path := "config/config.production.json"
			return path
		case "local":
			path := "config/config.local.json"
			return path
		case "alpha":
			path := "config/config.alpha.json"
			return path
		default:
			path := "config/config.production.json"
			return path
		}
	}
	path := getPath(enviroment)
	err := gonfig.GetConf(path, &Config)
	if err != nil {
		fmt.Println(err)
	}
}
