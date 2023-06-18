package dao

import (
	"github.com/HeisenbergAbhi/auction-service/internal/supply/dto"
	"github.com/HeisenbergAbhi/auction-service/internal/supply/storage"
)

// AdSpaceDAO provides the data access operations for ad spaces.
type AdSpaceDAO struct {
	storage *storage.AdSpaceStorage
}

// NewAdSpaceDAO creates a new AdSpaceDAO with the provided AdSpaceStorage.
func NewAdSpaceDAO(storage *storage.AdSpaceStorage) *AdSpaceDAO {
	return &AdSpaceDAO{
		storage: storage,
	}
}

// GetAllAdSpaces retrieves all ad spaces.
func (d *AdSpaceDAO) GetAllAdSpaces() ([]dto.AdSpaceDTO, error) {
	return d.storage.GetAllAdSpaces()
}

// GetAdSpaceByID retrieves an ad space by its ID.
func (d *AdSpaceDAO) GetAdSpaceByID(id int) (*dto.AdSpaceDTO, error) {
	return d.storage.GetAdSpaceByID(id)
}

// CreateAdSpace creates a new ad space and returns its ID.
func (d *AdSpaceDAO) CreateAdSpace(adSpace *dto.AdSpaceDTO) (int, error) {
	return d.storage.CreateAdSpace(adSpace)
}
