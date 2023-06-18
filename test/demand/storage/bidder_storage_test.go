package storage_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/HeisenbergAbhi/auction-service/internal/demand/dto"
	"github.com/HeisenbergAbhi/auction-service/internal/demand/storage"
)

func TestCreateBid(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO bids (bidder_id, ad_space_id, amount) VALUES (?, ?, ?)").
		WithArgs(1, 2, 15.0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	bidStorage := storage.NewBidStorage(db)
	bid := &dto.BidDTO{
		BidderID: 1,
		AdSpaceID: 2,
		Amount: 15.0,
	}

	err = bidStorage.CreateBid(bid)
	assert.NoError(t, err)
}

func TestGetBidsByAdSpace(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"bidder_id", "ad_space_id", "amount"}).
		AddRow(1, 2, 10.0).
		AddRow(3, 2, 15.0)

	mock.ExpectQuery("SELECT bidder_id, ad_space_id, amount FROM bids WHERE ad_space_id = ?").
		WithArgs(2).
		WillReturnRows(rows)

	bidStorage := storage.NewBidStorage(db)
	adSpaceID := 2

	result, err := bidStorage.GetBidsByAdSpace(adSpaceID)
	assert.NoError(t, err)

	expectedResult := []dto.BidDTO{
		{BidderID: 1, AdSpaceID: 2, Amount: 10.0},
		{BidderID: 3, AdSpaceID: 2, Amount: 15.0},
	}
	assert.Equal(t, expectedResult, result)
}

func TestGetWinningBidByAdSpace(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"bidder_id", "ad_space_id", "amount"}).
		AddRow(3, 2, 15.0)

	mock.ExpectQuery("SELECT bidder_id, ad_space_id, amount FROM bids WHERE ad_space_id = ? ORDER BY amount DESC LIMIT 1").
		WithArgs(2).
		WillReturnRows(rows)

	bidStorage := storage.NewBidStorage(db)
	adSpaceID := 2

	result, err := bidStorage.GetWinningBidByAdSpace(adSpaceID)
	assert.NoError(t, err)

	expectedResult := &dto.BidDTO{
		BidderID: 3, AdSpaceID: 2, Amount: 15.0,
	}
	assert.Equal(t, expectedResult, result)
}
