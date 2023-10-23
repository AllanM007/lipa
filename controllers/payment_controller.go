package controllers

import (
	"time"

	"github.com/AllanM007/lipa/models"
	"gorm.io/gorm"
)

type Payments struct {
	DB *gorm.DB
}

func PaymentsRepo(db *gorm.DB) *Payments {
	db.AutoMigrate(&models.PaymentChannel{})

	return &Payments{DB :db}
}

type Channel struct {
	Id              uint           `json:"id"`
	Identifier      string         `json:"identifier"`
	Name            string         `json:"name"`
}

type Currency struct {
	Id              uint           `json:"id"`
	Name            string         `json:"name"`
	Active          bool           `json:"active"`
}

type Payment struct {
	UID             string         `json:"uid"        binding:"required"`
	Source          Channel        `json:"source"     binding:"required"`
	Amount          float64        `json:"amount"     binding:"amount"`
	Currency        Currency       `json:"currency"   binding:"currency"`
	Recipient       Channel        `json:"recipient"  binding:"recipient"`
	Timestamp       time.Time      `json:"timeStamp"  binding:"timestamp"`
	Remarks         string         `json:"remarks"    binding:"remarks"`
}