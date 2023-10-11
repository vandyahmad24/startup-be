package handler

import (
	"fmt"
	"net/http"
	"startup/campaign"
	"startup/helper"
	"startup/logger"
	"startup/users"
	"strconv"

	"github.com/gin-gonic/gin"
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
	fmt.Println(userId)
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

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		h.logger.LogFatal("Error to get campaign", err)
		response := helper.ApiResponse(http.StatusBadRequest, nil, err.Error(), "error")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getCampaign, err := h.service.GetCampaignById(input)
	if err != nil {
		h.logger.LogFatal("Error to get campaign", err)
		response := helper.ApiResponse(http.StatusBadRequest, nil, "failed to get campaign", "error")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse(http.StatusOK, campaign.FormatCampaignDetail(getCampaign), "Get campaign", "success")
	c.JSON(http.StatusOK, response)
	return

}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		h.logger.LogFatal("CreateCampaignInput bind request", err)

		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse(http.StatusBadRequest, errorMessage, "Campaign Create Error", "error input campaign")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	current := c.MustGet("currentUser").(*users.User)
	input.User = *current

	campaignResponse, err := h.service.CreateCampaign(input)
	if err != nil {
		h.logger.LogFatal("CreateCampaign Create", err)
		response := helper.ApiResponse(http.StatusBadRequest, err.Error(), "Register Account Failed", "error register create")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse(http.StatusOK, campaignResponse, "Campaign has been created", "success")
	c.JSON(http.StatusOK, response)
}
