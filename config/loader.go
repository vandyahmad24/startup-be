package config

import (
	"fmt"
	"os"
	configApss "startup/config/config"
	"startup/config/dbconfig"
	"startup/util"

	"github.com/spf13/viper"
)

type config struct {
	Database dbconfig.DatabaseList
	Config   configApss.ConfigList
}

var cfg config

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(dir + "/config/dbconfig")
	viper.SetConfigType("yaml")
	viper.SetConfigName("database.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load database config: %v", err))
	}
	viper.Unmarshal(&cfg)

	viper.AddConfigPath(dir + "/config/config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load config: %v", err))
	}
	viper.Unmarshal(&cfg)

	fmt.Println("=============================")
	fmt.Println(util.Stringify(cfg))
	fmt.Println("=============================")
}

func GetConfig() *config {
	return &cfg
}
