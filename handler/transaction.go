package handler

import (
	"log"
	"net/http"
	"startup/helper"
	"startup/logger"
	"startup/payment"
	"startup/transaction"
	"startup/users"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service        transaction.Service
	servicePayment payment.Service
	logger         *logger.Logger
}

func NewTransactionHandler(service transaction.Service, servicePayment payment.Service) *transactionHandler {
	return &transactionHandler{service: service, servicePayment: servicePayment}
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

		response := helper.ApiResponse(http.StatusNotFound, nil, err.Error(), "error")
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.ApiResponse(http.StatusOK, transaction.FormatTransactionFormatter(transactionRes), "Create Transaction", "success")
	c.JSON(http.StatusOK, response)
	return

}

func (h *transactionHandler) GetNotification(c *gin.Context) {
	var input transaction.TransactionData
	err := c.ShouldBindJSON(&input)
	if err != nil {
		log.Println("error ", err)
		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse(http.StatusBadRequest, errorMessage, "GetNotification Failed", "error input register")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.ProcessPayment(input)
	if err != nil {
		response := helper.ApiResponse(http.StatusNotFound, nil, err.Error(), "error")
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.ApiResponse(http.StatusOK, nil, "Create Transaction", "success")
	c.JSON(http.StatusOK, response)
	return

}
