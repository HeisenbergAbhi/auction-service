package dao

import (
	"github.com/HeisenbergAbhi/auction-service/internal/demand/dto"
	"github.com/HeisenbergAbhi/auction-service/internal/demand/storage"
)

// BidderDAO provides the data access operations for bidders.
type BidderDAO struct {
	storage *storage.BidderStorage
}

// NewBidderDAO creates a new BidderDAO with the provided BidderStorage.
func NewBidderDAO(storage *storage.BidderStorage) *BidderDAO {
	return &BidderDAO{
		storage: storage,
	}
}

// GetAllBidders retrieves all bidders.
func (d *BidderDAO) GetAllBidders() ([]dto.BidderDTO, error) {
	return d.storage.GetAllBidders()
}

// GetBidderByID retrieves a bidder by their ID.
func (d *BidderDAO) GetBidderByID(id int) (*dto.BidderDTO, error) {
	return d.storage.GetBidderByID(id)
}

// CreateBidder creates a new bidder and returns their ID.
func (d *BidderDAO) CreateBidder(bidder *dto.BidderDTO) (int, error) {
	return d.storage.CreateBidder(bidder)
}
