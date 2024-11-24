package db

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DSN = "./gorm.db"
var DB *gorm.DB

func Dbconnection() {
	var error error
	DB, error = gorm.Open(sqlite.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Printf("DB connection established")
	}

}
