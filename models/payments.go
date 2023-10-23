package models

import "gorm.io/gorm"

type PaymentChannel struct {
	gorm.Model
	Name            string            `gorn:"name;unique"`
	Active          bool              `gorm:"active"`
}

func CreatePaymentChannel(db *gorm.DB, paymentChannel *PaymentChannel) (err error) {
	err = db.Create(paymentChannel).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPaymentChannels(db *gorm.DB, paymentChannel *[]PaymentChannel) (err error) {
	err = db.Order("created_at desc").Find(paymentChannel).Error
	if err != nil {
		return err
	}
	return nil
}