package handlers

import (
	"net/http"

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
	database.DB.Create(&issue)
	c.JSON(http.StatusCreated, issue)
}
