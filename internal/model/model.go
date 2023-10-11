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

type Drugs []string

type CreatePrescInput struct {
	PreID      int64  `json:"pre_id" gorm:"primaryKey;autoIncrement"`
	Patient    string `json:"patient"`
	Drugs      Drugs  `gorm:"type:text[]" json:"drugs"`
	Expiration string `json:"expiration"`
}

var Prescription CreatePrescInput

type JwtUser struct {
	Jwt      string "jwt"
	Password string "password"
	Role     string `json:"role"`
}
