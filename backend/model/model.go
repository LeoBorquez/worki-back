package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

// Init database
func Init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbHost, dbPort, username, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}
	db = conn
	db.Debug().AutoMigrate(&User{})

	//defer db.Close()
}

// returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}

// TestExport testing export
func TestExport() {
	fmt.Println("test export")
}
