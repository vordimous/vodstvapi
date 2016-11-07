package dao

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //import postgres
)

//DB ...
type DB struct {
	*sql.DB
}

const (
	//DbHost ...
	DbHost = "localdocker"
	//DbUser ...
	DbUser = "esvodsapi"
	//DbPassword ...
	DbPassword = "esvodsapi"
	//DbName ...
	DbName = "esvods"
)

var db *gorm.DB

//Init ...
func Init() {

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		DbUser, DbPassword, DbHost, DbName)

	var err error
	//set the global variable
	db, err = gorm.Open("postgres", dbinfo)
	checkErr(err, "db connect failed")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	err = db.DB().Ping()
	checkErr(err, "db ping failed")

	db.LogMode(true)
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}
