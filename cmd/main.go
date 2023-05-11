package main

import (
	"github.com/szmulinho/prescription/cmd/server"
	"github.com/szmulinho/prescription/database"
)

func main() {

	database.Connect()

	server.Run()
}
