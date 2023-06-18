package storage

import (
	"database/sql"

	"github.com/HeisenbergAbhi/auction-service/internal/supply/dto"
)

// AdSpaceStorage provides the storage operations for ad spaces.
type AdSpaceStorage struct {
	db *sql.DB
}

// NewAdSpaceStorage creates a new AdSpaceStorage with the provided MySQL database connection.
func NewAdSpaceStorage(db *sql.DB) *AdSpaceStorage {
	return &AdSpaceStorage{
		db: db,
	}
}

// GetAllAdSpaces retrieves all ad spaces.
func (s *AdSpaceStorage) GetAllAdSpaces() ([]dto.AdSpaceDTO, error) {
	rows, err := s.db.Query("SELECT id, name, base_price FROM ad_spaces")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	adSpaces := []dto.AdSpaceDTO{}
	for rows.Next() {
		var adSpace dto.AdSpaceDTO
		err := rows.Scan(&adSpace.ID, &adSpace.Name, &adSpace.BasePrice)
		if err != nil {
			return nil, err
		}
		adSpaces = append(adSpaces, adSpace)
	}

	return adSpaces, nil
}

// GetAdSpaceByID retrieves an ad space by its ID.
func (s *AdSpaceStorage) GetAdSpaceByID(id int) (*dto.AdSpaceDTO, error) {
	row := s.db.QueryRow("SELECT id, name, base_price FROM ad_spaces WHERE id = ?", id)

	var adSpace dto.AdSpaceDTO
	err := row.Scan(&adSpace.ID, &adSpace.Name, &adSpace.BasePrice)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &adSpace, nil
}

// CreateAdSpace creates a new ad space and returns its ID.
func (s *AdSpaceStorage) CreateAdSpace(adSpace *dto.AdSpaceDTO) (int, error) {
	result, err := s.db.Exec("INSERT INTO ad_spaces (name, base_price) VALUES (?, ?)", adSpace.Name, adSpace.BasePrice)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
