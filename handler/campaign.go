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
		response := helper.ApiResponse(http.StatusBadRequest, err.Error(), "Create campaign Failed", "error register create")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse(http.StatusOK, campaignResponse, "Campaign has been created", "success")
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputId campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&inputId)
	if err != nil {
		h.logger.LogFatal("Error to get campaign", err)
		response := helper.ApiResponse(http.StatusBadRequest, nil, err.Error(), "error")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&input)
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

	campaignResponse, err := h.service.UpdateCampaign(inputId, input)
	if err != nil {
		h.logger.LogFatal("UpdateCampgin Create", err)
		response := helper.ApiResponse(http.StatusBadRequest, err.Error(), "Update Campaign", "error update campaign")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse(http.StatusOK, campaignResponse, "Campaign has been update", "success")
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadCampaignImage(c *gin.Context) {
	var input campaign.CreateCampaingImage
	err := c.ShouldBind(&input)
	if err != nil {
		h.logger.LogFatal("UploadCampaignImage bind request", err)

		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse(http.StatusBadRequest, errorMessage, "Campaign Create Error", "error input campaign")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		h.logger.LogFatal("Upload avatar request", err)
		errorMessage := gin.H{"is_uploaded": false}
		response := helper.ApiResponse(http.StatusBadRequest, errorMessage, err.Error(), "error")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if file == nil {
		h.logger.LogFatal("Upload avatar request", err)
		errorMessage := gin.H{"is_uploaded": false}
		response := helper.ApiResponse(http.StatusBadRequest, errorMessage, "avatar is required", "error")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	current := c.MustGet("currentUser").(*users.User)
	// fmt.Println(current)
	userId := current.ID

	// path := "images/" + file.Filename
	path := fmt.Sprintf("upload/%d-%s", userId, file.Filename)
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

	resp, err := h.service.CreateCampaingImage(input, path)
	if err != nil {
		data := gin.H{"is_available": false}
		response := helper.ApiResponse(http.StatusBadRequest, data, err.Error(), "error")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse(http.StatusOK, resp, "success upload avatar", "success")
	c.JSON(http.StatusOK, response)

	return
}
