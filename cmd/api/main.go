package main

import (
	"io/ioutil"
	"log"
	"strconv"

	"github.com/PS-Sergey/proxy/internal/api"
	"gopkg.in/yaml.v2"
)

func main() {
	config := api.NewConfig()
	file, err := ioutil.ReadFile("config/api.yaml")
	if err != nil {
		log.Fatal("Can not find .yaml config file", err)
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Fatal("Can not marshall config file", err)
	}
	if _, err := strconv.Atoi(config.CacheSize); err != nil {
		log.Fatal("cacheSize must be number")
	}
	api := api.NewApi(config)
	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
}
