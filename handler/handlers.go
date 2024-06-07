package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sweetieSydney/postapi/database"
	"github.com/sweetieSydney/postapi/models"
)

func CreateIssue(c *gin.Context) {
	var issue models.Issue
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&issue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, issue)
}

func GetAllIssues(c *gin.Context) {
	var issues []models.Issue
	result := database.DB.Order("id asc").Find(&issues)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving issues"})
		return
	}

	c.JSON(http.StatusOK, issues)
}

func DeleteIssues(c *gin.Context) {
	ids := c.Query("ids")
	idSlice := strings.Split(ids, ",")

	var issues []models.Issue

	if err := database.DB.Where("id IN (?)", idSlice).Find(&issues).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Issue(s) not found"})
		return
	}

	if err := database.DB.Where("id IN (?)", idSlice).Delete(&models.Issue{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting issue(s)"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Issue(s) deleted successfully"})
}

func DeleteDuplicateIssues(c *gin.Context) {
	var issues []models.Issue
	database.DB.Find(&issues)

	seen := make(map[string]bool)
	var deletedIDs []uint
	for _, issue := range issues {
		if _, ok := seen[issue.Title]; ok {
			database.DB.Delete(&issue)
			deletedIDs = append(deletedIDs, issue.ID)
		} else {
			seen[issue.Title] = true
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Duplicate issues deleted successfully", "deleted_ids": deletedIDs})
}

func UpdateIssue(c *gin.Context) {
	id := c.Param("id")
	var issue models.Issue

	if err := database.DB.Where("id = ?", id).First(&issue).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Issue not found"})
		return
	}

	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := database.DB.Save(&issue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating issue"})
		return
	}

	c.JSON(http.StatusOK, issue)
}
