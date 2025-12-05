package model

import (
	"time"
)

type UserRegistration struct {
	ID           uint      `json:"id" gorm:"column:id;primarykey"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt    time.Time `json:"deleted_at" gorm:"column:deleted_at;index"`
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
	TID          string    `json:"tid" gorm:"column:tid"`
	MID          string    `json:"mid" gorm:"column:mid"`
	Status       string    `json:"status" gorm:"column:status"`
	Secret       string    `json:"secret" gorm:"column:secret"`
	Verification string    `json:"verification" gorm:"column:verification"`
}
