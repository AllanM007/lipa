package controllers

import (
	"time"

	"net/http"

	"github.com/AllanM007/lipa/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Payments struct {
	DB *gorm.DB
}

func PaymentsRepo(db *gorm.DB) *Payments {
	db.AutoMigrate(&models.PaymentChannel{})

	return &Payments{DB: db}
}

type Channel struct {
	Id         uint   `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

type Currency struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type Payment struct {
	UID       string    `json:"uid"        binding:"required"`
	Source    Channel   `json:"source"     binding:"required"`
	Amount    float64   `json:"amount"     binding:"amount"`
	Currency  Currency  `json:"currency"   binding:"currency"`
	Recipient Channel   `json:"recipient"  binding:"recipient"`
	Timestamp time.Time `json:"timeStamp"  binding:"timestamp"`
	Remarks   string    `json:"remarks"    binding:"remarks"`
}

func NewPayment(c *gin.Context) {
	var payment Payment

	if err := c.BindJSON(&payment); err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "ACCESS_DENIED", "message": "Invalid payment!!", "err": err.Error()})
		return
	}
}

type UpdatePay struct {
	Id            uint      `json:"id"         binding:"required"`
	Source        Channel   `json:"source"     binding:"required"`
	Amount        float64   `json:"amount"     binding:"amount"`
	Currency      Currency  `json:"currency"   binding:"currency"`
	Recipient     Channel   `json:"recipient"  binding:"recipient"`
	Timestamp     time.Time `json:"timeStamp"  binding:"timestamp"`
	UpdateRemarks string    `json:"updateRemarks"   binding:"required"`
}

func UpdatePayment(c *gin.Context) {

	var updatePay UpdatePay

	if err := c.BindJSON(&updatePay); err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "ACCESS_DENIED", "message": "Invalid payment update!!", "err": err.Error()})
		return
	}
}

type DeletePay struct {
	Id            uint   `json:"id" binding:"required"`
	DeleteRemarks string `json:"deleteRemarks"  binding:"required"`
}

func DeletePayment(c *gin.Context) {

}
