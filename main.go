package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	cfg "startup/config"
	"startup/handler"
	"startup/helper"
	"startup/logger"
	"startup/users"
	"syscall"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("halo")
	config := cfg.GetConfig()
	logger := logger.NewLogger()
	fmt.Println(config.Database.Startup.Mysql)
	dsn := "root:@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.LogFatal("error koneksi database", err.Error())
	}
	logger.LogInfo("success connect to database")
	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	router.NoRoute(func(ctx *gin.Context) {
		response := helper.ApiResponse(http.StatusNotFound, nil, "Route not found", "error route not found")
		ctx.JSON(404, response)
	})
	api := router.Group("/api/v1")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/check-email", userHandler.CheckEmail)
	api.POST("/upload-avatar", userHandler.UploadAvatar)
	go func() {
		router.Run(":8000")
	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v\n", signal.String())

}

// func Handler(c *gin.Context) {
// 	logger := logger.NewLogger()

// 	var users []users.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)

// }
