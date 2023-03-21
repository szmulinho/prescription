package endpoints

import (
	"encoding/json"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
)

func GetAllPrescriptions(w http.ResponseWriter, r *http.Request) {
	prescs := []model.Presc{}
	json.NewEncoder(w).Encode(prescs)

}
