package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open(sqlite.Open("lots.db"), &gorm.Config{
		CreateBatchSize: 1000,
	})
	check(err)

	db.AutoMigrate(&Lot{})
}
