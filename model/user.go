package model

import "time"

type User struct {
	Id             string `db:"id"`
	Name           string
	Email          string
	PhoneNumber    string `db:"phone_number"`
	Password       string
	Address        string
	BirthDate      time.Time `db:"birth_date"`
	ProfilePicture string    `db:"profile_picture"`
	CreatedAt      time.Time `db:"created_at"`
	UpdateAt       time.Time `db:"update_at"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Name        string    `json:"name" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	Address     string    `json:"address" binding:"required"`
	BirthDate   time.Time `json:"birth_date" time_format:"2006-01-02"`
}

type UserUpdate struct {
	Id          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type CheckEmail struct {
	Email string `json:"email" binding:"required,email"`
}
