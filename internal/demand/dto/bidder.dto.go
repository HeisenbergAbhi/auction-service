package dto

// BidderDTO represents the data transfer object for a bidder.
type BidderDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
