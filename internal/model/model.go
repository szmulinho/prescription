package model

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"os"
	"strconv"
)

var Prescs []CreatePrescInput

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type Response struct {
	Data string `json:"data"`
}

type StringArray []string

func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}
	var buf bytes.Buffer
	buf.WriteString("{")
	for i, s := range a {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(strconv.Quote(s))
	}
	buf.WriteString("}")
	return buf.String(), nil
}

func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal string array value: %v", value)
	}
	if len(s) == 0 {
		*a = make(StringArray, 0)
		return nil
	}
	s = bytes.Trim(s, "{}")
	parts := bytes.Split(s, []byte(","))
	res := make(StringArray, len(parts))
	for i, p := range parts {
		unquoted, err := strconv.Unquote(string(bytes.TrimSpace(p)))
		if err != nil {
			return fmt.Errorf("failed to unmarshal string array value: %v", value)
		}
		res[i] = unquoted
	}
	*a = res
	return nil
}

type CreatePrescInput struct {
	PreId      string      `json:"preid"`
	Drugs      StringArray `gorm:"type:text[]" json:"drugs"`
	Expiration string      `json:"expiration"`
}

var Prescription CreatePrescInput

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Login    string `gorm:"unique"`
	Password string
}

type JwtUser struct {
	Jwt      string "jwt"
	Password string "password"
}
