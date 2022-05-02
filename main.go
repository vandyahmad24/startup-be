package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"startup/auth"
	"startup/campaign"
	cfg "startup/config"
	"startup/handler"
	"startup/helper"
	"startup/logger"
	"startup/middleware"
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
	// fmt.Println()
	confDB := config.Database.Startup.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", confDB.Username, confDB.Password, confDB.Host, confDB.Port, confDB.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
		logger.LogFatal("error koneksi database", err.Error())
	}
	logger.LogInfo("success connect to database")
	// deklrasi middleware

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	authService := auth.NewService()
	authMiddleware := middleware.NewAuthMiddleware(userService, authService)
	userHandler := handler.NewUserHandler(userService, authService)

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	router := gin.Default()
	router.Static("/images", "./images")
	router.NoRoute(func(ctx *gin.Context) {
		response := helper.ApiResponse(http.StatusNotFound, nil, "Route not found", "error route not found")
		ctx.JSON(404, response)
	})
	api := router.Group("/api/v1")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/check-email", userHandler.CheckEmail)
	api.POST("/upload-avatar", authMiddleware.AuthMiddleware, userHandler.UploadAvatar)
	//campaign
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
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
