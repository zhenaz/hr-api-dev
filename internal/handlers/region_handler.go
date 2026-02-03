package handlers

import (
	"net/http"
	"strconv"

	"codeid.hr-api/internal/models"
	"codeid.hr-api/internal/services"
	"github.com/gin-gonic/gin"
)

type RegionHandler struct {
	regionService services.RegionService
}

func NewRegionHandler(regionService services.RegionService) *RegionHandler {
	return &RegionHandler{
		regionService: regionService,
	}
}
func (h *RegionHandler) GetRegions(c *gin.Context) {
	regions, err := h.regionService.GetAllRegions(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to fetch regions",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  regions,
	})
}
func (h *RegionHandler) GetRegion(c *gin.Context) {
	// get id param from endpoint router yg bertipe string
	//stcconv.ParUint : konvert string ke uint
	//parseUint(idstr,10,32) -> 10 base ten number,32 tipe bit
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid region ID",
		})
		return
	}
	region, err := h.regionService.GetRegionByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": "Region not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    region,
	})
}
func (h *RegionHandler) CreateRegion(c *gin.Context) {
	var region models.Region
	if err := c.ShouldBindJSON(&region); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid input",
			"details": err.Error(),
		})
		return
	}
	if err := h.regionService.CreateRegion(c.Request.Context(), &region); err !=
		nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to create region",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Region created successfully",
		"data":    region,
	})
}
func (h *RegionHandler) UpdateRegion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid region ID",
		})
		return
	}
	var region models.Region
	if err := c.ShouldBindJSON(&region); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid input",
			"details": err.Error(),
		})
		return
	}
	region.RegionID = uint(id)
	if err := h.regionService.UpdateRegion(c.Request.Context(), &region); err !=
		nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to update region",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Region updated successfully",
		"data":    region,
	})
}
func (h *RegionHandler) DeleteRegion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid region ID",
		})
		return
	}
	if err := h.regionService.DeleteRegion(c.Request.Context(), uint(id)); err !=
		nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to delete region",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Region deleted successfully",
	})
}

func (h *RegionHandler) GetRegionsWithCountry(c *gin.Context) {
	data, err := h.regionService.GetRegionsWithCountries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *RegionHandler) GetRegionByIdWithCountry(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	data, err := h.regionService.GetRegionWithCountries(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
		return
	}

	c.JSON(http.StatusOK, data)
}

