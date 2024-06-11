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
	//This function will bind the request body to the issue struct
	//and return an error if the request body is not a valid JSON
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//This function will create a new issue in the database
	if err := database.DB.Create(&issue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, issue)
}

func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if username is more than 32 characters
	if len(newUser.Username) > 32 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot be more than 32 characters"})
		return
	}

	// Save newUser to your database
	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "user created"})

}

func GetIssue(c *gin.Context) {
	var issue models.Issue
	id := c.Param("id")
	//This function will retrieve the issue from the database
	if err := database.DB.Where("id = ?", id).First(&issue).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Issue not found"})
		return
	}

	c.JSON(http.StatusOK, issue)
}

func GetUser(c *gin.Context) {
	var user models.User
	username := c.Param("username")
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetAllIssues(c *gin.Context) {
	var issues []models.Issue
	//This function will retrieve all issues from the database
	//and order them by id in ascending order as previously they did not order in order
	result := database.DB.Order("id asc").Find(&issues)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving issues"})
		return
	}

	c.JSON(http.StatusOK, issues)
}

func DeleteIssues(c *gin.Context) {
	//This function will retrieve the ids of the issues to be deleted
	//Then it will split the ids into a slice of strings
	ids := c.Query("ids")
	idSlice := strings.Split(ids, ",")

	var issues []models.Issue
	//This function will retrieve the issues from the database
	if err := database.DB.Where("id IN (?)", idSlice).Find(&issues).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Issue(s) not found"})
		return
	}
	//This function will delete the issues from the database
	if err := database.DB.Where("id IN (?)", idSlice).Delete(&models.Issue{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting issue(s)"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Issue(s) deleted successfully"})
}

func DeleteDuplicateIssues(c *gin.Context) {
	var issues []models.Issue
	database.DB.Find(&issues)
	//seen is a map that will store the title of the issues as keys
	seen := make(map[string]bool)
	var deletedIDs []uint

	//This loop will iterate through the issues and delete the duplicate issues
	for _, issue := range issues {
		if _, ok := seen[issue.Title]; ok {
			//This function will delete the duplicate issues from the database
			database.DB.Delete(&issue)
			//This will append the ids of the deleted issues to the deletedIDs slice
			deletedIDs = append(deletedIDs, issue.ID)
		} else {
			seen[issue.Title] = true
		}
	}
	//This function will return a JSON response with the message and the ids of the deleted issues
	c.JSON(http.StatusOK, gin.H{"message": "Duplicate issues deleted successfully", "deleted_ids": deletedIDs})
}

func UpdateIssue(c *gin.Context) {
	id := c.Param("id")
	var issue models.Issue
	//This function will retrieve the issue from the database
	if err := database.DB.Where("id = ?", id).First(&issue).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Issue not found"})
		return
	}
	//This function will bind the request body to the issue struct
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	//This function will update the issue in the database
	if err := database.DB.Save(&issue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating issue"})
		return
	}

	c.JSON(http.StatusOK, issue)
}
