package http

import (
	"auth-server/internal/adapters/dto"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"auth-server/internal/application/ports/incoming"
)

type AuthHandler struct {
	authService incoming.AuthPort
}

func NewAuthHandler(authService incoming.AuthPort) *AuthHandler {
	return &AuthHandler{authService: authService}
}

//func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
//	var user dto.LoginRequestDTO
//	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	token, err := h.authService.Login(user)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusUnauthorized)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(token)
//}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInputDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Error while decoding register request: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := h.authService.Register(user)
	if err.Message != "" {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, output)
		fmt.Println("User created successfully")
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
