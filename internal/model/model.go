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

type Response struct {
	Data string `json:"data"`
}

type CreatePrescInput struct {
	PreId      string   `json:"pre-id"`
	Drugs      []string `json:"drugs"`
	Expiration string   `json:"expiration"`
}

var Prescription CreatePrescInput

type Drug struct {
	DrugID string `json:"drugID"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Price  int    `json:"price"`
}

var Drugs = []Drug{
	{
		DrugID: "1",
		Name:   "Apap",
		Type:   "Painkiller",
		Price:  10,
	},
	{
		DrugID: "2",
		Name:   "Ibuprom",
		Type:   "Painkiller",
		Price:  8,
	},
	{
		DrugID: "3",
		Name:   "Magnefar",
		Type:   "Suplement",
		Price:  100,
	},
}

type User struct {
	UserID   string `json:"user-id"`
	Password string `json:"password"`
}

var Users = []User{
	User{
		UserID:   "user1",
		Password: "password1",
	},
	User{
		UserID:   "user2",
		Password: "password2",
	},
}
