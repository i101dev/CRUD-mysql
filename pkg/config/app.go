package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {

	dsn := "user:userpass@tcp(localhost:3306)/mydb?charset=utf8&parseTime=True&loc=Local"

	d, err := gorm.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	} else {
		db = d
		fmt.Println("\nMySQL connectin success")
	}
}

func GetDB() *gorm.DB {
	return db
}
