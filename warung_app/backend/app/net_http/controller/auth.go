package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cocoasterr/warung_app/domain"
	"github.com/cocoasterr/warung_app/service"
	"github.com/google/uuid"
)

type AuthController struct {
	Service service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{
		Service: service,
	}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed!", http.StatusInternalServerError)
		return
	}
	var payload domain.User
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	validEmail, _ := c.Service.CheckEmail(payload.Email)
	if !validEmail {
		http.Error(w, "Email wrong!", http.StatusBadRequest)
		return
	}
	if len(payload.Password) < 6 {
		http.Error(w, "Password must at least 6 characther", http.StatusBadRequest)
		return
	}
	existUser := c.Service.CheckUser(payload, payload.TbName())
	if existUser != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	} // var err error
	payload.Id = uuid.New().String()
	payload.Password, _ = c.Service.HashPassword(payload.Password)
	// if err != nil {
	// 	http.Error(w, "Failed Hashed Password", http.StatusInternalServerError)
	// 	return
	// }
	if err := c.Service.Create(payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := map[string]interface{}{"Message": "Success!"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
