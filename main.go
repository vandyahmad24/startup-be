package main

import (
	"fmt"
	cfg "startup/config"
	"startup/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("halo")
	config := cfg.GetConfig()
	logger := logger.NewLogger()
	fmt.Println(config.Database.Startup.Mysql)

	dsn := "root:@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.LogFatal(err.Error())
	}
	logger.LogInfo("success connect to database")

}
