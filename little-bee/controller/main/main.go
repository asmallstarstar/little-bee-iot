package main

import (
	"controller/maininit"
	"os"
	"util/casbin"
	"util/log"
	yaml "util/yaml/service"
)

func main() {
	maininit.Initialization()
	if !maininit.MQTTInit() {
		log.Logger.Error("failed to init mqtt,exit")
		os.Exit(0)
	}
	casbin.CasbinInit()
	maininit.SetWebRunMode(yaml.Yaml.Service.WebRunMode)
	r := maininit.GinEngine()
	r.Run(yaml.Yaml.Service.Host)
}
