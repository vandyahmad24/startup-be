package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"startup/helper"
	"startup/logger"
	"startup/transaction"
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
