package main

import (
	"fmt"
	"os"

	"github.com/darkit/godis/config"
	"github.com/darkit/godis/lib/logger"
	"github.com/darkit/godis/lib/utils"
	"github.com/darkit/godis/redis/server"
	"github.com/darkit/godis/tcp"
)

var banner = `
	   ______          ___
	  / ____/___  ____/ (_)____
	 / / __/ __ \/ __  / / ___/
	/ /_/ / /_/ / /_/ / (__  )
	\____/\____/\__,_/_/____/

`

var defaultProperties = &config.ServerProperties{
	Bind:           "0.0.0.0",
	Port:           6399,
	AppendOnly:     false,
	AppendFilename: "",
	MaxClients:     1000,
	RunID:          utils.RandString(40),
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

func main() {
	print(banner)
	configFilename := os.Getenv("CONFIG")
	if configFilename == "" {
		if fileExists("redis.conf") {
			config.SetupConfig("redis.conf")
		} else {
			config.Properties = defaultProperties
		}
	} else {
		config.SetupConfig(configFilename)
	}
	err := tcp.ListenAndServeWithSignal(&tcp.Config{
		Address: fmt.Sprintf("%s:%d", config.Properties.Bind, config.Properties.Port),
	}, server.MakeHandler())
	if err != nil {
		logger.Error(err.Error())
	}
}
