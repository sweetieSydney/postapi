package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sweetieSydney/postapi/database"
	"github.com/sweetieSydney/postapi/routes"
)

type Issue struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Labels      []string `json:"labels"`
	Assignees   []string `json:"assignees"`
}

func main() {
	database.Connect()

	router := routes.SetupRouter(database.DB)
	router.Run(":8080")
}
