package service

import (
	"errors"

	"github.com/HeisenbergAbhi/auction-service/internal/demand/dao"
	"github.com/HeisenbergAbhi/auction-service/internal/demand/dto"
)

// BidderService provides the business logic operations for bidders.
type BidderService struct {
	dao *dao.BidderDAO
}

// NewBidderService creates a new BidderService with the provided BidderDAO.
func NewBidderService(dao *dao.BidderDAO) *BidderService {
	return &BidderService{
		dao: dao,
	}
}

// GetAllBidders retrieves all bidders.
func (s *BidderService) GetAllBidders() ([]dto.BidderDTO, error) {
	return s.dao.GetAllBidders()
}

// GetBidderByID retrieves a bidder by their ID.
func (s *BidderService) GetBidderByID(id int) (*dto.BidderDTO, error) {
	if id <= 0 {
		return nil, errors.New("invalid bidder ID")
	}
	return s.dao.GetBidderByID(id)
}

// CreateBidder creates a new bidder.
func (s *BidderService) CreateBidder(bidder *dto.BidderDTO) (int, error) {
	if bidder == nil || bidder.Name == "" || bidder.Email == "" {
		return 0, errors.New("invalid bidder data")
	}
	return s.dao.CreateBidder(bidder)
}
