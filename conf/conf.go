package conf

import (
	"log"
	"os"
	"gopkg.in/ini.v1"
	"strconv"
)

type conf map[string]string

func (c conf) Get(key string, defVal ...string) string {
	if val, exist := c[key]; exist {
		return val
	}
	if len(defVal) > 0 {
		return defVal[0]
	}
	return ""
}

func (c conf) GetI(key string, defVal ...int) int {
	if val := c.Get(key); val != "" {
		if val, err := strconv.Atoi(val); err != nil {
			panic(err)
		} else {
			return val
		}
	}
	if len(defVal) > 0 {
		return defVal[0]
	}
	return 0
}

func GetI(key string, defVal ...int) int {
	return C.GetI(key, defVal...)
}

var C = make(conf, 10)

func init() {
	load()
}

const (
	configName = "/app.conf"
)

func load() {
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
		C[key.Name()] = key.Value()
	}
	log.Println(C)
}
