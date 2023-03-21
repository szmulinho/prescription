package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/api/endpoints"
	"github.com/szmulinho/prescription/internal/api/jwt"
	"log"
	"net/http"
)

func Run() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Starting the application...")
	router.HandleFunc("/", endpoints.HomeLink)
	router.HandleFunc("/presc", endpoints.CreatePrescription).Methods("POST")
	router.HandleFunc("/presc/{id}", endpoints.GetOnePrescription).Methods("GET")
	router.HandleFunc("/prescs", jwt.ValidateMiddleware(endpoints.GetAllPrescriptions)).Methods("GET")
	router.HandleFunc("/presc/{id}", endpoints.UpdatePrescription).Methods("PATCH")
	router.HandleFunc("/presc/{id}", endpoints.DeletePrescription).Methods("DELETE")
	router.HandleFunc("/authenticate", jwt.CreateToken).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), router))
}

func server() {
	Run()
}
