package model

import (
	"github.com/lib/pq"
	"os"
)

type Prescription struct {
	PreID      int64          `json:"pre_id" gorm:"primaryKey;autoIncrement"`
	Drugs      pq.StringArray `gorm:"type:text[]" json:"drugs" sql:"type:text[]"`
	Patient    string         `json:"patient"`
	Expiration string         `json:"expiration"`
}

var Presc Prescription

var Prescriptions []Prescription

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type JwtUser struct {
	Jwt      string "jwt"
	Password string "password"
	Role     string `json:"role"`
}
