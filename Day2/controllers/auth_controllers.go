package controllers

import (
	"Day2/middleware"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var login LoginRequest
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// Dummy check (replace with real DB check)
	if login.Username != "ahmad@gmail.com" || login.Password != "yaad12345" {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	token, err := middleware.GenerateJWT(login.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
