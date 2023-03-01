package model

import "time"

type User struct {
	Id             string
	Name           string
	Email          string `json:"email" binding:"required,email"`
	PhoneNumber    string `db:"phone_number"`
	Password       string `json:"password" binding:"required"`
	Address        string
	BirthDate      time.Time `db:"birth_date"`
	ProfilePicture string    `db:"profile_picture"`
	CreatedAt      time.Time `db:"created_at"`
	UpdateAt       time.Time `db:"update_at"`
}

type UserLogin struct {
	Email    string
	Password string
}
