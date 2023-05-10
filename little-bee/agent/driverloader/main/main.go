package main

import (
	dao "dao/database"
	loader "driverloader/driverloader"
	"fmt"
	"os"
	"os/signal"
	"util/log"
	yaml "util/yaml/agent"
)

func main() {
	fmt.Println("driver loader")
	yaml.Init()
	log.InitLog(
		yaml.Yaml.Log.DriverLoaderName,
		yaml.Yaml.Log.MaxAge,
		yaml.Yaml.Log.MaxBackups,
		yaml.Yaml.Log.MaxAge,
		yaml.Yaml.Log.Level)
	dao.InitDB(yaml.Yaml.Database.Dsn, yaml.Yaml.Database.LogLevel)

	loader.DriverManager.InitGrpc()
	loader.DriverManager.LoadDriver()
	loader.DriverManager.CreateDriverType()
	loader.DriverManager.CreateDriver()

	//CTRL+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("got signal:", s)

	loader.DriverManager.Stop()
}
