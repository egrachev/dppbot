package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url   string
	Title string
}

func initDB() {
	var err error

	db, err = gorm.Open(sqlite.Open(linksDB), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Link{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
