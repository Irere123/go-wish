package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Wish struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey; not null"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

func ConnectDb() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&Wish{})

	return
}
