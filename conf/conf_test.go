package conf

import (
	"testing"
	"log"
)

func TestConf_Get(t *testing.T) {
	log.Println(C.Get("a") == "a")
	//log.Println(C.Get("c") == "c")

	log.Println(C.GetI("b") == 1)
	//log.Println(C.GetI("c") == 2)

	log.Println(GetI("c") == 2)
}
