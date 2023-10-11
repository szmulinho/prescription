package model

import (
	"os"
)

var Prescriptions []Prescription

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type Drugs []string

type Prescription struct {
	PreID      int64  `json:"pre_id" gorm:"primaryKey;autoIncrement"`
	Drugs      Drugs  `gorm:"type:text[]" json:"drugs"`
	Patient    string `json:"patient"`
	Expiration string `json:"expiration"`
}

var Presc Prescription

type JwtUser struct {
	Jwt      string "jwt"
	Password string "password"
	Role     string `json:"role"`
}
