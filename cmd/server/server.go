package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/api/endpoints"
	"github.com/szmulinho/prescription/internal/api/endpoints/prescriptions/add"
	"github.com/szmulinho/prescription/internal/api/endpoints/prescriptions/delete"
	"github.com/szmulinho/prescription/internal/api/endpoints/prescriptions/get"
	"github.com/szmulinho/prescription/internal/api/endpoints/prescriptions/update"
	"github.com/szmulinho/prescription/internal/api/endpoints/users/login"
	"github.com/szmulinho/prescription/internal/api/endpoints/users/register"
	"github.com/szmulinho/prescription/internal/api/endpoints/users/userData"
	"github.com/szmulinho/prescription/internal/api/jwt"
	"log"
	"net/http"
)

func Run() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Starting the application...")
	router.HandleFunc("/", endpoints.HomeLink)
	router.HandleFunc("/presc", add.CreatePrescription).Methods("POST")
	router.HandleFunc("/presc/{id}", get.GetOnePrescription).Methods("GET")
	router.HandleFunc("/prescs", (get.GetAllPrescriptions)).Methods("GET")
	router.HandleFunc("/presc/{id}", update.UpdatePrescription).Methods("PATCH")
	router.HandleFunc("/presc/{id}", delete.DeletePrescription).Methods("DELETE")
	router.HandleFunc("/authenticate", jwt.CreateToken).Methods("POST")
	router.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		userID := uint(1)
		isDoctor := true
		token, err := jwt.GenerateToken(w, r, int64(userID), isDoctor)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(token))
	}).Methods("POST")
	router.HandleFunc("/login", login.Login).Methods("POST")
	router.HandleFunc("/register", register.CreateUser).Methods("POST")
	router.HandleFunc("/user", userData.GetUserDataHandler).Methods("GET")
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"}),
		handlers.ExposedHeaders([]string{}),
		handlers.AllowCredentials(),
		handlers.MaxAge(86400),
	)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), cors(router)))
}

func server() {
	Run()
}
