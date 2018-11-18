package config

import (
	"github.com/alive2212/go-illuminate/config"
)

var	ExposePort = base.GetEnv("EXPOSE_PORT", "10000")

// SQLHost host of mongodb to connect
var SQLHost = base.GetEnv("SQL_HOST", "127.0.0.1")

// SQLPort is port of mysql
var SQLPort = base.GetEnv("SQL_PORT", "3306")

// SQLUserName is sql username
var SQLUserName = base.GetEnv("SQL_USER_NAME", "gouser")

// SQLPassword is sql password
var SQLPassword = base.GetEnv("SQL_PASSWORD", "gopassword")

// SQLDBName name of current project database
var SQLDBName = base.GetEnv("SQL_DB_NAME", "GoTest")

// database driver
var DBDriver = base.GetEnv("DB_DRIVER","mysql")

func New() {

}



