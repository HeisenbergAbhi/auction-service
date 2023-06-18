package storage

import (
	"database/sql"
	"errors"

	"github.com/HeisenbergAbhi/auction-service/internal/demand/dto"
)

// BidderStorage provides the storage implementation for bidders.
type BidderStorage struct {
	db *sql.DB
}

// NewBidderStorage creates a new BidderStorage with the provided database connection.
func NewBidderStorage(db *sql.DB) *BidderStorage {
	return &BidderStorage{
		db: db,
	}
}

// GetAllBidders retrieves all bidders from the storage.
func (s *BidderStorage) GetAllBidders() ([]dto.BidderDTO, error) {
	rows, err := s.db.Query("SELECT id, name, email FROM bidders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bidders := make([]dto.BidderDTO, 0)
	for rows.Next() {
		var bidder dto.BidderDTO
		if err := rows.Scan(&bidder.ID, &bidder.Name, &bidder.Email); err != nil {
			return nil, err
		}
		bidders = append(bidders, bidder)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bidders, nil
}

// GetBidderByID retrieves a bidder by their ID from the storage.
func (s *BidderStorage) GetBidderByID(id int) (*dto.BidderDTO, error) {
	row := s.db.QueryRow("SELECT id, name, email FROM bidders WHERE id = ?", id)

	var bidder dto.BidderDTO
	err := row.Scan(&bidder.ID, &bidder.Name, &bidder.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &bidder, nil
}

// CreateBidder creates a new bidder in the storage and returns their ID.
func (s *BidderStorage) CreateBidder(bidder *dto.BidderDTO) (int, error) {
	result, err := s.db.Exec("INSERT INTO bidders (name, email) VALUES (?, ?)", bidder.Name, bidder.Email)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}