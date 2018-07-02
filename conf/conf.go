package conf

import (
	"log"
	"os"
	"gopkg.in/ini.v1"
	"strconv"
)

type conf map[string]string

func (c conf) Get(key string) string {
	if val, exist := c[key]; exist {
		return val
	}
	panic("No value for key : '" + key + "'")
}

func (c conf) GetI(key string) int {
	val, err := strconv.Atoi(c.Get(key))
	if err != nil {
		panic(err)
	}
	return val
}

func GetI(key string) int {
	return C.GetI(key)
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
