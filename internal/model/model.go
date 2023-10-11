package model

import (
	"os"
)

var Prescs []CreatePrescInput

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type CreatePrescInput struct {
	PreID      int64  `json:"preid" gorm:"primaryKey;autoIncrement"`
	Patient    string `json:"patient"`
	Drugs      string `json:"drugs" gorm:"many2many:prescription_drugs;"`
	Expiration string `json:"expiration"`
}

var Prescription CreatePrescInput

type JwtUser struct {
	Jwt      string "jwt"
	Password string "password"
	Role     string `json:"role"`
}
