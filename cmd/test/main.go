package main

import (
	"github.com/brunetto/figa"
	"github.com/brunetto/goutils/conf"
	"log"
)

func main () {
	var (
		f = figa.FApp{}
		err error
		fileName string = "figa.json"
	)
	err = conf.LoadJsonConf(fileName, &f)
	if err != nil {
		log.Fatal("Error reading JSON config file: ", err)
	}
	log.Println(f)
}