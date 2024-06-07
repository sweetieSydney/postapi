package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sweetieSydney/postapi/database"
	"github.com/sweetieSydney/postapi/routes"
	"gorm.io/gorm"
)

type Issue struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Labels      []string `json:"labels"`
	Assignees   []string `json:"assignees"`
}

var DB *gorm.DB

// func init() {
// 	var err error
// 	DB, err = gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic(fmt.Sprintf("failed to connect database: %v", err))
// 	}
// }

func main() {
	database.Connect()

	// var err error
	// DB, err = gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic(fmt.Sprintf("failed to connect database: %v", err))
	// }
	//routes.SetupRoutes(database.DB)
	router := routes.SetupRouter(database.DB)
	router.Run(":8080")
}
