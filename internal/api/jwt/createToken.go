package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/szmulinho/prescription/internal/model"
	"net/http"
	"time"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserID,
		"password": user.Password,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})
	tokenString, error := token.SignedString(model.JwtKey)
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(model.JwtToken{Token: tokenString})
}
