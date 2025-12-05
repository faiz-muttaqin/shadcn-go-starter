package model

import (
	"time"
)

type UserRejected struct {
	ID           uint      `json:"id" gorm:"column:id;primarykey"`
	RegisterID   uint      `json:"regist_id" gorm:"column:regist_id"`
	RegisteredAt time.Time `json:"registered_at" gorm:"column:registered_at"`
	RejectedAt   time.Time `json:"rejected_at" gorm:"column:rejected_at"`
	RejectedBy   uint      `json:"rejected_by" gorm:"column:rejected_by"`
	FirstName    string    `json:"first_name" gorm:"column:first_name"`
	LastName     string    `json:"last_name" gorm:"column:last_name"`
	UserName     string    `json:"user_name" gorm:"column:user_name"`
	Email        string    `json:"email" gorm:"column:email"`
	Phone        string    `json:"phone" gorm:"column:phone"`
	Password     string    `json:"password" gorm:"column:password"`
	PlaceOfBirth string    `json:"place_of_birth" gorm:"column:place_of_birth"`
	DateOfBirth  string    `json:"date_of_birth" gorm:"column:date_of_birth"`
	Country      string    `json:"country" gorm:"column:country"`
	Province     string    `json:"province" gorm:"column:province"`
	District     string    `json:"district" gorm:"column:district"`
	Address      string    `json:"address" gorm:"column:address"`
	PostalCode   string    `json:"postal_code" gorm:"column:postal_code"`
	IDCardImage  string    `json:"id_card_image" gorm:"column:id_card_image"`
	UserImage    string    `json:"user_image" gorm:"column:user_image"`
	Latitude     string    `json:"latitude" gorm:"column:latitude"`
	Longitude    string    `json:"longitude" gorm:"column:longitude"`
	IccId        string    `json:"IccId" gorm:"column:IccId"`
	IMEI         string    `json:"imei" gorm:"column:imei"`
}