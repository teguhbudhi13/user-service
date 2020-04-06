package model

import (
	"github.com/jinzhu/gorm"
	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// User schema
type User struct {
	ID   int    `gorm:"primary_key; auto_increment" sql:"id" json:"id"`
	Name string `gorm:"not null" json:"name"`
	DOB  string `gorm:"type:date; default:null" json:"dob"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
