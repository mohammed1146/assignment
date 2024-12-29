package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mohammed1146/skelton/internal/domain"
	"github.com/mohammed1146/skelton/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.User true "User details"
// @Success 201 {object} map[string]string
// @Router /register [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.service.Register(ctx, &user); err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// Login godoc
// @Summary Login user
// @Description Authenticate user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "User credentials"
// @Success 200 {object} map[string]string
// @Router /login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	token, err := h.service.Login(ctx, credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
