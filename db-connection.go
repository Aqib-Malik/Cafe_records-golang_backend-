package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

// var usrIDSN = "server = localhost; port = 3306; user = root; password = ; database = testdb"

var usrIDSN = "root:@tcp(localhost:3306)/testdb?parseTime=true"
var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(usrIDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("Connection Failed..:(")
	}

	Database.AutoMigrate(&Cafe{})

}
