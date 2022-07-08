package main

import (
	"time"

	"gorm.io/gorm"
)

type Lot struct {
	gorm.Model
	Lot               string `gorm:"uniqueIndex"`
	Item              string
	PurchaseOrder     string
	OnHandQuantity    int
	SterilizationDate time.Time
	UseExpirationDate time.Time
}

func (l *Lot) getFromDatabase() *Lot {
	if err := db.Where("lot = ?", l.Lot).First(&l).Error; err != nil {
		panic(err)
	}
	return l
}
