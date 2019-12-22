package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

const (
	kDBUsername string = "STRANGER_BOT_1_DB_USERNAME"
	kDBPassword string = "STRANGER_BOT_1_DB_PASSWORD"
	kDBHost     string = "STRANGER_BOT_1_DB_HOST"
	kDBName     string = "STRANGER_BOT_1_DB_NAME"
)

func init() {
	e := godotenv.Load()
	if e != nil {
		panic(e)
	}

	username := os.Getenv(kDBUsername)
	password := os.Getenv(kDBPassword)
	host := os.Getenv(kDBHost)
	dbName := os.Getenv(kDBName)

	dbUri := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbName)

	conn, err := gorm.Open("mysql", dbUri)

	if err != nil {
		panic(err)
	}

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
