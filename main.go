package main

import (
	"fmt"
	"net/http"
	cfg "startup/config"
	"startup/logger"
	"startup/users"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("halo")
	config := cfg.GetConfig()
	logger := logger.NewLogger()
	fmt.Println(config.Database.Startup.Mysql)
	logger.LogInfo("aplikasi dijalankan")
	router := gin.Default()
	router.GET("/", Handler)
	router.Run()
}

func Handler(c *gin.Context) {
	logger := logger.NewLogger()
	dsn := "root:@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.LogFatal(err.Error())
	}
	logger.LogInfo("success connect to database")
	var users []users.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)

}
