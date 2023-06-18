package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/HeisenbergAbhi/auction-service/internal/supply/dto"
	"github.com/HeisenbergAbhi/auction-service/internal/supply/service"
)

// AdSpaceHandler handles the HTTP requests for ad spaces.
type AdSpaceHandler struct {
	service *service.AdSpaceService
}

// NewAdSpaceHandler creates a new AdSpaceHandler with the provided AdSpaceService.
func NewAdSpaceHandler(service *service.AdSpaceService) *AdSpaceHandler {
	return &AdSpaceHandler{
		service: service,
	}
}

// GetAllAdSpaces handles the request to get all ad spaces.
func (h *AdSpaceHandler) GetAllAdSpaces(w http.ResponseWriter, r *http.Request) {
	adSpaces, err := h.service.GetAllAdSpaces()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(adSpaces)
}

// GetAdSpaceByID handles the request to get an ad space by its ID.
func (h *AdSpaceHandler) GetAdSpaceByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Path[len("/adspaces/"):]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ad space ID", http.StatusBadRequest)
		return
	}

	adSpace, err := h.service.GetAdSpaceByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if adSpace == nil {
		http.Error(w, "Ad space not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(adSpace)
}

// CreateAdSpace handles the request to create a new ad space.
func (h *AdSpaceHandler) CreateAdSpace(w http.ResponseWriter, r *http.Request) {
	var adSpace dto.AdSpaceDTO
	err := json.NewDecoder(r.Body).Decode(&adSpace)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateAdSpace(&adSpace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"id": %d}`, id)
}
