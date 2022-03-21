package handler

import (
	"net/http"
	"startup/helper"
	"startup/logger"
	"startup/users"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService users.Service
	logger      *logger.Logger
}

func NewUserHandler(userService users.Service) *userHandler {
	logger := logger.NewLogger()
	return &userHandler{
		userService: userService,
		logger:      logger,
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
		response := helper.ApiResponse(http.StatusInternalServerError, err.Error(), "Register Account Failed", "error register create")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	formatter := users.FormatUser(newUser, "123")

	response := helper.ApiResponse(http.StatusOK, formatter, "Account has been created", "success")
	c.JSON(http.StatusOK, response)
}