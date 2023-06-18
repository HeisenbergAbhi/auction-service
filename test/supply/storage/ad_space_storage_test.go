package storage_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/HeisenbergAbhi/auction-service/internal/supply/dto"
	"github.com/HeisenbergAbhi/auction-service/internal/supply/storage"
)

func TestGetAllAdSpaces(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "base_price"}).
		AddRow(1, "Ad Space 1", 10.0).
		AddRow(2, "Ad Space 2", 15.0)

	mock.ExpectQuery("SELECT id, name, base_price FROM ad_spaces").WillReturnRows(rows)

	adSpaceStorage := storage.NewAdSpaceStorage(db)
	adSpaces, err := adSpaceStorage.GetAllAdSpaces()
	if err != nil {
		t.Fatal(err)
	}

	expectedAdSpaces := []dto.AdSpaceDTO{
		{ID: 1, Name: "Ad Space 1", BasePrice: 10.0},
		{ID: 2, Name: "Ad Space 2", BasePrice: 15.0},
	}

	assert.Equal(t, expectedAdSpaces, adSpaces)
}

func TestGetAdSpaceByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "base_price"}).
		AddRow(1, "Ad Space 1", 10.0)

	mock.ExpectQuery("SELECT id, name, base_price FROM ad_spaces WHERE id = ?").
		WithArgs(1).
		WillReturnRows(rows)

	adSpaceStorage := storage.NewAdSpaceStorage(db)
	adSpace, err := adSpaceStorage.GetAdSpaceByID(1)
	if err != nil {
		t.Fatal(err)
	}

	expectedAdSpace := &dto.AdSpaceDTO{ID: 1, Name: "Ad Space 1", BasePrice: 10.0}

	assert.Equal(t, expectedAdSpace, adSpace)
}

func TestCreateAdSpace(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO ad_spaces (name, base_price) VALUES (?, ?)").
		WithArgs("Ad Space 3", 20.0).
		WillReturnResult(sqlmock.NewResult(3, 1))

	adSpaceStorage := storage.NewAdSpaceStorage(db)
	adSpace := &dto.AdSpaceDTO{Name: "Ad Space 3", BasePrice: 20.0}
	id, err := adSpaceStorage.CreateAdSpace(adSpace)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3, id)
}
