package controllers

import (
	"time"

	"golang.org/x/exp/constraints"
)

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

func Add[T constraints.Ordered](a , b T) T {
	return a + b
}

func Subtract[T constraints.Ordered](a , b T) T {
	return a + b
}

func Multiply[T constraints.Ordered](a , b T) T {
	return a + b
}

func Divide[T constraints.Ordered](a , b T) T {
	return a + b
}

func Min[T constraints.Ordered](a , b T) T {
	if a > b {
		return b
	}
	return a
}

func Max[T constraints.Ordered](a , b T) T {
	if a > b {
		return a
	}
	return b
}