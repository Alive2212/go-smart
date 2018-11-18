package database

import (
	"../config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var  DB *gorm.DB

/**
Create db object as singleton design pattern
 */
func New() {
	var err error
	connectionString := GetDbString()
	DB, err = gorm.Open(
		config.DBDriver,
		connectionString)

	if err != nil {
		panic("failed to connect database")
	}
}

/**
Close DB Connection
 */
func DBClose() {
	defer DB.Close()
}

//GetDbString yo
func GetDbString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.SQLUserName, config.SQLPassword, config.SQLHost, config.SQLPort, config.SQLDBName)
}
