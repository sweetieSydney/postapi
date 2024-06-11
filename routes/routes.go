package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/sweetieSydney/postapi/handler"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.POST("/issues", handlers.CreateIssue)
	router.GET("/issues", handlers.GetAllIssues)
	router.DELETE("/issues", handlers.DeleteIssues)
	router.DELETE("/issues/duplicates", handlers.DeleteDuplicateIssues)
	router.PUT("/issues/:id", handlers.UpdateIssue)

	router.POST("/users", handlers.CreateUser)
	router.GET("/users/:username", handlers.GetUser)

	return router
}
