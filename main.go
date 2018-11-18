package main

import (
	"./bootstrap"
	"./config"
	"./database"
	"./routes"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func main() {
	bootstrap.New()
	config.New()
	database.New()
	routes.New()
	bootstrap.Start()
}
