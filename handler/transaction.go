package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"startup/helper"
	"startup/logger"
	"startup/transaction"
	"startup/users"
)

type transactionHandler struct {
	service transaction.Service
	logger  *logger.Logger
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service: service}
}
func (h *transactionHandler) GetCampaginTransaction(c *gin.Context) {
	var input transaction.GetTransactionCampaignInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		h.logger.LogFatal("Error to get campaign", err)
		response := helper.ApiResponse(http.StatusBadRequest, nil, err.Error(), "error")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactionRes, err := h.service.GetTransactionByCampaignId(input)
	if err != nil {
		h.logger.LogFatal("Error to get campaign", err)
		response := helper.ApiResponse(http.StatusNotFound, nil, "Error to get campaign", "error")
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.ApiResponse(http.StatusOK, transaction.FormatCampaignTransactions(transactionRes), "List of transaction", "success")
	c.JSON(http.StatusOK, response)
	return

}

func (h *transactionHandler) GetCampaginTransactionByUserId(c *gin.Context) {
	current := c.MustGet("currentUser").(*users.User)

	transactionRes, err := h.service.GetTransactionByUserId(current.ID)
	if err != nil {
		h.logger.LogFatal("Error to get campaign", err)
		response := helper.ApiResponse(http.StatusNotFound, nil, "Error to get campaign", "error")
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.ApiResponse(http.StatusOK, transaction.FormatUserTransactionFormatters(transactionRes), "List of transaction", "success")
	c.JSON(http.StatusOK, response)
	return

}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		h.logger.LogFatal("CreateTransaction bind request", err)

		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse(http.StatusBadRequest, errorMessage, "CreateTransaction Create Error", "error input CreateTransaction")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	current := c.MustGet("currentUser").(*users.User)
	input.User = *current

	transactionRes, err := h.service.CreateTransaction(input)
	if err != nil {
		h.logger.LogFatal("Error to get CreateTransaction", err)
		response := helper.ApiResponse(http.StatusNotFound, nil, "Error to get CreateTransaction", "error")
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.ApiResponse(http.StatusOK, transaction.FormatTransactionFormatter(transactionRes), "Create Transaction", "success")
	c.JSON(http.StatusOK, response)
	return

}
