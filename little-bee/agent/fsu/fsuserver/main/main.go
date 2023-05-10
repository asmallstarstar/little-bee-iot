package main

import (
	dao "dao/database"
	"fmt"
	f "fsuserver/fsumanager"
	"os"
	"os/signal"
	log "util/log"
	yaml "util/yaml/agent"
)

func main() {
	fmt.Println("fsu server")
	yaml.Init()
	log.InitLog(
		yaml.Yaml.Log.FsuFileName,
		yaml.Yaml.Log.MaxAge,
		yaml.Yaml.Log.MaxBackups,
		yaml.Yaml.Log.MaxAge,
		yaml.Yaml.Log.Level)
	dao.InitDB(yaml.Yaml.Database.Dsn, yaml.Yaml.Database.LogLevel)

	f.FsuManager.InitGrpc()
	f.FsuManager.RegistFsuReflectType()
	f.FsuManager.LoadFsu()
	f.FsuManager.CreateFsuByAgent()

	//CTRL+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("got signal:", s)

	f.FsuManager.Stop()
}
