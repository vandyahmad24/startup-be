package handler

import (
	"fmt"
	"net/http"
	"startup/auth"
	"startup/helper"
	"startup/logger"
	"startup/users"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService users.Service
	logger      *logger.Logger
	authService auth.Service
}

func NewUserHandler(userService users.Service, authService auth.Service) *userHandler {
	logger := logger.NewLogger()
	return &userHandler{
		userService: userService,
		logger:      logger,
		authService: authService,
	}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input users.RegisterInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		h.logger.LogFatal("RegisterUser bind request", err)

		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse(http.StatusBadRequest, errorMessage, "Register Account Failed", "error input register")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUser, err := h.userService.Register(&input)
	if err != nil {
		h.logger.LogFatal("RegisterUser Create", err)
		response := helper.ApiResponse(http.StatusBadRequest, err.Error(), "Register Account Failed", "error register create")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		h.logger.LogFatal("RegisterUser Token", err)
		response := helper.ApiResponse(http.StatusBadRequest, err.Error(), "Register Token Failed", "error register create")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := users.FormatUser(newUser, token)

	response := helper.ApiResponse(http.StatusOK, formatter, "Account has been created", "success")
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input users.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		h.logger.LogFatal("LoginUser bind request", err)

		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse(http.StatusBadRequest, errorMessage, "Login Failed", "error input login")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	user, err := h.userService.Login(&input)
	if err != nil {
		h.logger.LogFatal("Login", err)
		response := helper.ApiResponse(http.StatusBadRequest, err.Error(), "LoginFailed", "error login")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		h.logger.LogFatal("LoginFailed Token", err)
		response := helper.ApiResponse(http.StatusBadRequest, err.Error(), "Login Token Failed", "error login create")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := users.FormatUser(user, token)

	response := helper.ApiResponse(http.StatusOK, formatter, "Login Succesfully", "success")
	c.JSON(http.StatusOK, response)
	return
}

func (h *userHandler) CheckEmail(c *gin.Context) {
	var input users.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		h.logger.LogFatal("CheckEmail bind request", err)
		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse(http.StatusBadRequest, errorMessage, "Check email Failed", "error check email")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.userService.CheckEmail(&input)
	if err != nil {
		data := gin.H{"is_available": false}
		response := helper.ApiResponse(http.StatusBadRequest, data, "Check Email", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_available": true}
	response := helper.ApiResponse(http.StatusOK, data, "Email Available", "email available")
	c.JSON(http.StatusOK, response)
	return
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		h.logger.LogFatal("Upload avatar request", err)
		errorMessage := gin.H{"is_uploaded": false}
		response := helper.ApiResponse(http.StatusBadRequest, errorMessage, err.Error(), "error")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userId := 22
	// path := "images/" + file.Filename
	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)
	// fmt.Println(path)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		h.logger.LogFatal("Upload avatar request", err)
		errorMessage := gin.H{"is_uploaded": false}
		response := helper.ApiResponse(http.StatusBadRequest, errorMessage, err.Error(), "error")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//sementara

	_, err = h.userService.SaveAvatar(userId, path)
	if err != nil {
		data := gin.H{"is_available": false}
		response := helper.ApiResponse(http.StatusBadRequest, data, err.Error(), "error")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_available": true}
	response := helper.ApiResponse(http.StatusOK, data, "success upload avatar", "success")
	c.JSON(http.StatusOK, response)

	return
}
