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
	"startup/payment"
	"startup/transaction"
	"startup/users"
	"syscall"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config := cfg.NewConfig()
	logger := logger.NewLogger()
	// fmt.Println()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
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

	transactionRepository := transaction.NewRepository(db)

	paymentService := payment.NewService(campaignRepository)

	transactionService := transaction.NewService(transactionRepository, paymentService, campaignRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService, paymentService)

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Static("/upload", "./upload")
	router.NoRoute(func(ctx *gin.Context) {
		response := helper.ApiResponse(http.StatusNotFound, nil, "Route not found", "error route not found")
		ctx.JSON(404, response)
	})
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, "Welcome To Startup Server Running in Port 8181")
	})
	api := router.Group("/api/v1")
	api.Use(authMiddleware.LoggingMiddleware)
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/check-email", userHandler.CheckEmail)
	api.POST("/upload-avatar", authMiddleware.AuthMiddleware, userHandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleware.AuthMiddleware, userHandler.FetchUser)
	//campaign
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware.AuthMiddleware, campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware.AuthMiddleware, campaignHandler.UpdateCampaign)
	api.POST("/campaigns-image", authMiddleware.AuthMiddleware, campaignHandler.UploadCampaignImage)

	api.GET("/campaigns/:id/transaction", authMiddleware.AuthMiddleware, transactionHandler.GetCampaginTransaction)
	api.GET("/transactions", authMiddleware.AuthMiddleware, transactionHandler.GetCampaginTransactionByUserId)
	api.POST("/transactions", authMiddleware.AuthMiddleware, transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)
	go func() {
		router.Run(fmt.Sprintf(":%s", config.PORT))
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
