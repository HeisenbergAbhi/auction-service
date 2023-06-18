package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/HeisenbergAbhi/auction-service/internal/demand/dto"
	"github.com/HeisenbergAbhi/auction-service/internal/demand/service"
)

// BidderHandler handles HTTP requests for bidders.
type BidderHandler struct {
	service *service.BidderService
}

// NewBidderHandler creates a new BidderHandler with the provided BidderService.
func NewBidderHandler(service *service.BidderService) *BidderHandler {
	return &BidderHandler{
		service: service,
	}
}

// GetAllBidders handles the request to retrieve all bidders.
func (h *BidderHandler) GetAllBidders(w http.ResponseWriter, r *http.Request) {
	bidders, err := h.service.GetAllBidders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bidders)
}

// GetBidderByID handles the request to retrieve a bidder by their ID.
func (h *BidderHandler) GetBidderByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Path[len("/bidders/"):]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid bidder ID", http.StatusBadRequest)
		return
	}

	bidder, err := h.service.GetBidderByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if bidder == nil {
		http.Error(w, "Bidder not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bidder)
}

// CreateBidder handles the request to create a new bidder.
func (h *BidderHandler) CreateBidder(w http.ResponseWriter, r *http.Request) {
	var bidder dto.BidderDTO
	err := json.NewDecoder(r.Body).Decode(&bidder)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateBidder(&bidder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"id": %d}`, id)
}
