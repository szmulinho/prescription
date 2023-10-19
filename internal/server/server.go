package server

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/szmulinho/prescription/internal/server/endpoints"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Run(ctx context.Context, db *gorm.DB) {
	handler := endpoints.NewHandler(db)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/presc", handler.AddPrescription).Methods("POST")
	router.HandleFunc("/presc/{id}", handler.GetOnePrescription).Methods("GET")
	router.HandleFunc("/patient/{patient}", handler.GetPrescriptionsForPatient).Methods("GET")
	router.HandleFunc("/prescs", (handler.GetAllPrescriptions)).Methods("GET")
	router.HandleFunc("/presc/{id}", handler.UpdatePrescription).Methods("PATCH")
	router.HandleFunc("/presc/{id}", handler.DeletePrescription).Methods("DELETE")
	router.HandleFunc("/authenticate", handler.CreateToken).Methods("POST")
	router.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		userID := uint(1)
		isDoctor := true
		token, err := handler.GenerateToken(w, r, int64(userID), isDoctor)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(token))
	}).Methods("POST")
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"}),
		handlers.ExposedHeaders([]string{}),
		handlers.AllowCredentials(),
		handlers.MaxAge(86400),
	)
	go func() {
		err := http.ListenAndServe(":8080", cors(router))
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
}
