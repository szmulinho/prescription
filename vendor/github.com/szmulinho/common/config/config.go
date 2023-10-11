package config

import "fmt"

type StorageConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     string `json:"port"`
	Sslmode  string `json:"sslmode"`
}

func (c StorageConfig) ConnectionString() string {
	return fmt.Sprintf("host=%spostgres user=%spostgres password=%sL96a1prosniper dbname=%sszmul-med port=%s5432 sslmode=%sdisable", c.Host, c.User, c.Password, c.Dbname, c.Port, c.Sslmode)
}
