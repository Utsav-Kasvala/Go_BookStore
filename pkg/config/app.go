package config

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql" 
)
var (db *gorm.DB)

// SetupDB initializes the database instance

func Connect(){
	d,err := gorm.Open("mysql","utsav:utsav123@utsav/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db = d

}

// GetDB returns a handle to the DB object

func GetDB() *gorm.DB {
	return db
}



