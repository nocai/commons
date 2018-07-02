package conf

import (
	"log"
	"os"
	"gopkg.in/ini.v1"
)

func init() {
	initConfig()
}

var Config = make(map[string]interface{}, 10)

const (
	configName = "/app.conf"
)

func initConfig() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	log.Printf(pwd + configName)
	f, err := ini.Load(pwd + configName)
	if err != nil {
		panic(err)
	}

	section, err := f.GetSection("DEFAULT")
	if err != nil {
		panic(err)
	}
	keys := section.Keys()
	for _, key := range keys {
		Config[key.Name()] = key.Value()
	}
	log.Println(Config)
}
