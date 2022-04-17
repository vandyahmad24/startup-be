package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"startup/campaign"
	"startup/helper"
	"startup/logger"
	"strconv"
)

type campaignHandler struct {
	service campaign.Service
	logger  *logger.Logger
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	logger := logger.NewLogger()
	return &campaignHandler{service: service, logger: logger}
}
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.FindCampaigns(userId)
	if err != nil {
		h.logger.LogFatal("Error to get campaign", err)
		response := helper.ApiResponse(http.StatusNotFound, nil, "Error to get campaign", "error")
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := helper.ApiResponse(http.StatusOK, campaign.FormatCampaigns(campaigns), "List of campaign", "success")
	c.JSON(http.StatusOK, response)
	return
}
