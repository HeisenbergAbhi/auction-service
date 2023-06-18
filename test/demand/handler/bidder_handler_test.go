package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/HeisenbergAbhi/auction-service/internal/demand/dto"
	"github.com/HeisenbergAbhi/auction-service/internal/demand/handler"
	"github.com/HeisenbergAbhi/auction-service/internal/demand/service"
)

func TestGetAllBidsByAdSpace(t *testing.T) {
	bids := []dto.BidDTO{
		{BidderID: 1, AdSpaceID: 2, Amount: 10.0},
		{BidderID: 3, AdSpaceID: 2, Amount: 15.0},
	}

	bidService := &service.BidServiceMock{
		GetBidsByAdSpaceFunc: func(adSpaceID int) ([]dto.BidDTO, error) {
			return bids, nil
		},
	}

	handler := handler.NewBidHandler(bidService)

	req, err := http.NewRequest("GET", "/bids/adspaces/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler.GetAllBidsByAdSpace(recorder, req)

	expectedResponse := `[{"bidder_id":1,"ad_space_id":2,"amount":10},{"bidder_id":3,"ad_space_id":2,"amount":15}]`
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedResponse, recorder.Body.String())
}

func TestGetWinningBidByAdSpace(t *testing.T) {
	winningBid := &dto.BidDTO{BidderID: 3, AdSpaceID: 2, Amount: 15.0}

	bidService := &service.BidServiceMock{
		GetWinningBidByAdSpaceFunc: func(adSpaceID int) (*dto.BidDTO, error) {
			return winningBid, nil
		},
	}

	handler := handler.NewBidHandler(bidService)

	req, err := http.NewRequest("GET", "/bids/adspaces/2/winning", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler.GetWinningBidByAdSpace(recorder, req)

	expectedResponse := `{"bidder_id":3,"ad_space_id":2,"amount":15}`
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedResponse, recorder.Body.String())
}

func TestPlaceBid(t *testing.T) {
	bid := &dto.BidDTO{BidderID: 1, AdSpaceID: 2, Amount: 10.0}

	bidService := &service.BidServiceMock{
		PlaceBidFunc: func(bid *dto.BidDTO) error {
			return nil
		},
	}

	handler := handler.NewBidHandler(bidService)

	payload, err := json.Marshal(bid)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/bids/place", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler.PlaceBid(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}
