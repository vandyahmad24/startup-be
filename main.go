package main

import (
	"fmt"
	cfg "startup/config"
	"startup/logger"
)

func main() {
	fmt.Println("halo")
	config := cfg.GetConfig()
	logger := logger.InitLogrus()
	logger.Info("INI info")
	fmt.Println(config.Database.Startup.Mysql)

}
