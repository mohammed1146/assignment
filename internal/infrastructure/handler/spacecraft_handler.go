package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mohammed1146/skelton/internal/service"
)

type SpacecraftHandler struct {
	service service.SpacecraftService
}

func NewSpacecraftHandler(service service.SpacecraftService) *SpacecraftHandler {
	return &SpacecraftHandler{service: service}
}

// ListSpacecrafts godoc
// @Summary List spacecrafts
// @Description Get a list of users with optional name or class or status filtering
// @Tags spacecrafts
// @Produce json
// @Param name query string false "Spacecraft name filter"
// @Param class query string false "Spacecraft class filter"
// @Param status query string false "Spacecraft status filter"
// @Success 200 {array} domain.Spacecraft
// @Router /spacecrafts [get]
func (h *SpacecraftHandler) ListSpacecrafts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	name := r.URL.Query().Get("name")
	class := r.URL.Query().Get("class")
	status := r.URL.Query().Get("status")

	spacecrafts, err := h.service.ListSpacecrafts(ctx, name, class, status)
	fmt.Println(spacecrafts)
	if err != nil {
		http.Error(w, "Failed to fetch spacecrafts", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(spacecrafts)
}
