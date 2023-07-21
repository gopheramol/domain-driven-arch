package handler

import (
	"net/http"
	"strconv"

	"github.com/gopheramol/domain-driven-arch/internal/ad/service"
	"github.com/gopheramol/domain-driven-arch/pkg/db/orm/sqlc"

	"github.com/gin-gonic/gin"
)

// AdHandler handles HTTP requests related to ads.
type AdHandler struct {
	adService service.AdService
}

// NewAdHandler creates a new instance of AdHandler.
func NewAdHandler(adService service.AdService) *AdHandler {
	return &AdHandler{
		adService: adService,
	}
}

// CreateAd handles the HTTP POST request to create a new ad.
func (h *AdHandler) CreateAd(c *gin.Context) {
	var ad sqlc.Ad
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ad, err := h.adService.CreateAd(ad)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ad)
}

// GetAllAds handles the HTTP GET request to retrieve all ads.
func (h *AdHandler) GetAllAds(c *gin.Context) {
	ads, err := h.adService.GetAllAds()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ads)
}

// GetAdByID handles the HTTP GET request to retrieve an ad by its ID.
func (h *AdHandler) GetAdByID(c *gin.Context) {
	adID := c.Param("id")

	adIDInt, _ := strconv.Atoi(adID)

	ad, err := h.adService.GetAdByID(int64(adIDInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ad)
}

// UpdateAd handles the HTTP PUT request to update an existing ad.
func (h *AdHandler) UpdateAd(c *gin.Context) {
	adID := c.Param("id")

	var updatedAd sqlc.Ad
	if err := c.ShouldBindJSON(&updatedAd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	adIDInt, _ := strconv.Atoi(adID)

	err := h.adService.UpdateAd(int64(adIDInt), updatedAd)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ad updated successfully"})
}

// DeleteAd handles the HTTP DELETE request to delete an ad.
func (h *AdHandler) DeleteAd(c *gin.Context) {
	adID := c.Param("id")
	adIDInt, _ := strconv.Atoi(adID)

	err := h.adService.DeleteAd(int64(adIDInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ad deleted successfully"})
}

// GetAdByUserID handles the HTTP GET request to retrieve ads by a specific user ID.
func (h *AdHandler) GetAdByUserID(c *gin.Context) {
	userID := c.Param("user_id")

	userIDInt, _ := strconv.Atoi(userID)

	userAds, err := h.adService.GetAdByUserID(int64(userIDInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userAds)
}
