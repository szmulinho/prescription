package login

import (
	"encoding/json"
	"github.com/szmulinho/prescription/internal/api/jwt"
	"github.com/szmulinho/prescription/internal/database"
	"github.com/szmulinho/prescription/internal/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user model.User
	result := database.DB.Where("login = ?", credentials.Login).First(&user)
	if result.Error != nil {
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
		return
	}

	var isDoctor bool

	if user.Role == "doctor" {
		isDoctor = true
	}

	token, err := jwt.GenerateToken(w, r, uint(user.ID), isDoctor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
	w.Write(responseJSON)
}
