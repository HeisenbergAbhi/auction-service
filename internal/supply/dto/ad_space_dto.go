package dto

// AdSpaceDTO represents the data transfer object for an ad space.
type AdSpaceDTO struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	BasePrice  float64 `json:"base_price"`
}
