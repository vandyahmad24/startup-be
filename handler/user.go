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
		h.logger.LogFatal(err)
		// log.Println(err)
		response := helper.ApiResponse(http.StatusBadRequest, nil, "Register Account Failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUser, err := h.userService.Register(&input)
	if err != nil {
		// h.logger.LogFatal(err)
		response := helper.ApiResponse(http.StatusInternalServerError, nil, "Register Account Failed", "error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	formatter := users.FormatUser(newUser, "123")

	response := helper.ApiResponse(http.StatusOK, formatter, "Account has been created", "success")
	c.JSON(http.StatusOK, response)
}
