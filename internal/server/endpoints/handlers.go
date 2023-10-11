package endpoints

import (
	"gorm.io/gorm"
	"net/http"
)

type Handlers interface {
	CreatePrescription(w http.ResponseWriter, r *http.Request)
	CreateToken(w http.ResponseWriter, r *http.Request)
	DeletePrescription(w http.ResponseWriter, r *http.Request)
	GenerateToken(w http.ResponseWriter, r *http.Request, userID int64, isDoctor bool) (string, error)
	GetAllPrescriptions(w http.ResponseWriter, r *http.Request)
	GetPrescriptionsForPatient(w http.ResponseWriter, r *http.Request)
	GetOnePrescription(w http.ResponseWriter, r *http.Request)
	UpdatePrescription(w http.ResponseWriter, r *http.Request)
	ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc
}

type handlers struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) Handlers {
	return &handlers{
		db: db,
	}
}
