package database

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MYSQL Driver
)

//DataBase :
type DataBase struct {
	*gorm.DB
}

//Configure : Configure the database connection
func Configure() (*DataBase, error) {
	hostname := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	database := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PWD")

	dbURL := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + database + "?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", dbURL)
	db.LogMode(true)
	return &DataBase{DB: db}, err
}

//CreateAndConnect :
func CreateAndConnect() (*DataBase, error) {

	hostname := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	database := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PWD")

	dbURL := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/"
	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}

	result := db.Exec("CREATE DATABASE " + database)
	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Exec("USE " + database)
	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Exec("CREATE TABLE users ( id integer, email varchar(255), name varchar(255), password varchar(255) )")
	if result.Error != nil {
		panic(result.Error)
	}

	db.LogMode(true)
	return &DataBase{DB: db}, err
}
