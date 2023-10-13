package model

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"os"
)

type Drugs []string

type Prescription struct {
	PreID      int64  `json:"pre_id" gorm:"primaryKey;autoIncrement"`
	Drugs      Drugs  `gorm:"type:text[]" json:"drugs"`
	Patient    string `json:"patient"`
	Expiration string `json:"expiration"`
}

func (d *Drugs) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), d)
}

func (d Drugs) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (p *Prescription) BeforeCreate(tx *gorm.DB) error {
	// Convert Drugs slice to PostgreSQL array
	p.Drugs = Drugs{"Apap", "AnotherDrug"} // Example drugs
	return nil
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
