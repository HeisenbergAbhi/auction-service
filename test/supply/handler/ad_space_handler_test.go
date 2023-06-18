package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/HeisenbergAbhi/auction-service/internal/supply/dto"
	"github.com/HeisenbergAbhi/auction-service/internal/supply/handler"
	"github.com/HeisenbergAbhi/auction-service/internal/supply/service"
)

func TestGetAllAdSpaces(t *testing.T) {
	adSpaces := []dto.AdSpaceDTO{
		{ID: 1, Name: "Ad Space 1", BasePrice: 10.0},
		{ID: 2, Name: "Ad Space 2", BasePrice: 15.0},
	}

	adSpaceService := &service.AdSpaceServiceMock{
		GetAllAdSpacesFunc: func() ([]dto.AdSpaceDTO, error) {
			return adSpaces, nil
		},
	}

	handler := handler.NewAdSpaceHandler(adSpaceService)

	req, err := http.NewRequest("GET", "/adspaces", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler.GetAllAdSpaces(recorder, req)

	expectedResponse := `[{"id":1,"name":"Ad Space 1","base_price":10},{"id":2,"name":"Ad Space 2","base_price":15}]`
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedResponse, recorder.Body.String())
}

func TestGetAdSpaceByID(t *testing.T) {
	adSpace := &dto.AdSpaceDTO{ID: 1, Name: "Ad Space 1", BasePrice: 10.0}

	adSpaceService := &service.AdSpaceServiceMock{
		GetAdSpaceByIDFunc: func(id int) (*dto.AdSpaceDTO, error) {
			return adSpace, nil
		},
	}

	handler := handler.NewAdSpaceHandler(adSpaceService)

	req, err := http.NewRequest("GET", "/adspaces/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler.GetAdSpaceByID(recorder, req)

	expectedResponse := `{"id":1,"name":"Ad Space 1","base_price":10}`
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedResponse, recorder.Body.String())
}

func TestCreateAdSpace(t *testing.T) {
	adSpace := &dto.AdSpaceDTO{Name: "Ad Space 3", BasePrice: 20.0}

	adSpaceService := &service.AdSpaceServiceMock{
		CreateAdSpaceFunc: func(adSpace *dto.AdSpaceDTO) (int, error) {
			return 3, nil
		},
	}

	handler := handler.NewAdSpaceHandler(adSpaceService)

	payload, err := json.Marshal(adSpace)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/adspaces/create", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler.CreateAdSpace(recorder, req)

	expectedResponse := `{"id":3}`
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedResponse, recorder.Body.String())
}
