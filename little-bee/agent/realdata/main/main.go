package main

import (
	dao "dao/database"
	"fmt"
	"os"
	"os/signal"
	r "realdata/realdataserver"
	log "util/log"
	yaml "util/yaml/agent"
)

func main() {
	fmt.Println("realdata server")
	yaml.Init()
	log.InitLog(
		yaml.Yaml.Log.RealdataFileName,
		yaml.Yaml.Log.MaxAge,
		yaml.Yaml.Log.MaxBackups,
		yaml.Yaml.Log.MaxAge,
		yaml.Yaml.Log.Level)
	dao.InitDB(yaml.Yaml.Database.Dsn, yaml.Yaml.Database.LogLevel)

	if !r.RealdataServerManager.InitMQTT() {
		log.Logger.Error("failed to init mqtt,exit")
		os.Exit(0)
	}

	if !r.RealdataServerManager.LoadSignal() {
		log.Logger.Error("failed to load config,exit")
		os.Exit(0)
	}
	r.RealdataServerManager.InitSignalStatus()
	r.RealdataServerManager.InitGrpc()

	//CTRL+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("got signal:", s)

	r.RealdataServerManager.StopServer()
}
