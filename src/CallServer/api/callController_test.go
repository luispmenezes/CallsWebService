package api

import (
	"CallServer/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type CallMockPersistenceManager struct {
}

func (m *CallMockPersistenceManager) AddCalls(calls *[]model.Call) error {
	if calls == nil || len(*calls) == 0 {
		return errors.New("invalid call list")
	}
	if (*calls)[0].Caller != "123" || (*calls)[0].IsInbound != true || (*calls)[0].Duration != 0 {
		return errors.New("call marshaling error")
	}

	return nil
}

func (m *CallMockPersistenceManager) RemoveCall(filterParams map[string]interface{}) (int,error) {
	if len(filterParams) != 3 {
		return 0,errors.New("invalid params")
	}
	return 1,nil
}

func (m *CallMockPersistenceManager) GetCalls(filterParams map[string]interface{}, pageIdx, pageSize int) (model.CallQueryResult, error) {
	if len(filterParams) != 3 {
		return model.CallQueryResult{}, errors.New("invalid params")
	}
	return model.CallQueryResult{
		Page:       pageIdx,
		TotalPages: pageIdx,
		PageSize:   pageSize,
		Result: []model.Call{{
			Caller:    "123",
			Callee:    "321",
			StartTime: time.Time{},
			EndTime:   time.Time{},
			IsInbound: false,
			Duration:  20,
			Cost:      30,
		}, {
			Caller:    "456",
			Callee:    "789",
			StartTime: time.Time{},
			EndTime:   time.Time{},
			IsInbound: true,
			Duration:  40,
			Cost:      50,
		}},
	}, nil
}

func (m *CallMockPersistenceManager) GetMetadata() ([]model.CallMetadata, error) {
	panic("not implemented")
}

func performCallRequest(r http.Handler, body io.Reader, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	return rec
}

func TestBaseController_CreateCalls_EmptyBody(t *testing.T) {
	baseController := NewBaseController(&CallMockPersistenceManager{})
	baseController.initializeRoutes()

	response := performCallRequest(baseController.Engine, http.NoBody, "PUT", "/call")

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestBaseController_CreateCalls(t *testing.T) {
	baseController := NewBaseController(&CallMockPersistenceManager{})
	baseController.initializeRoutes()

	callList := []model.Call{{
		Caller:    "123",
		Callee:    "321",
		StartTime: time.Time{},
		EndTime:   time.Time{},
		IsInbound: true,
		Duration:  20,
		Cost:      30,
	}}

	jsonBody, err := json.Marshal(callList)
	if err != nil {
		t.Fatal(err)
	}

	response := performCallRequest(baseController.Engine, bytes.NewBuffer(jsonBody), "PUT", "/call")

	assert.Equal(t, http.StatusCreated, response.Code)
}

func TestBaseController_RemoveCall(t *testing.T) {
	baseController := NewBaseController(&CallMockPersistenceManager{})
	baseController.initializeRoutes()

	caller := "123"
	callee := "321"
	startTime := time.Date(2020, time.January, 1, 1, 1, 1, 1, time.UTC)

	response := performCallRequest(baseController.Engine, http.NoBody, "DELETE",
		fmt.Sprintf("/call?caller=%s&callee=%s&startTime=%s", caller, callee, startTime.Format(time.RFC3339)))

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestBaseController_GetAllCalls(t *testing.T) {
	baseController := NewBaseController(&CallMockPersistenceManager{})
	baseController.initializeRoutes()

	caller := "123"
	callee := "321"
	startTime := time.Date(2020, time.January, 1, 1, 1, 1, 1, time.UTC)

	response := performCallRequest(baseController.Engine, http.NoBody, "GET",
		fmt.Sprintf("/call?caller=%s&callee=%s&startTime=%s", caller, callee, startTime.Format(time.RFC3339)))

	assert.Equal(t, http.StatusOK, response.Code)

	var callQueryResponse model.CallQueryResult
	err := json.Unmarshal([]byte(response.Body.String()), &callQueryResponse)

	assert.Nil(t, err)
	assert.Equal(t, callQueryResponse.Page, 0)
	assert.Equal(t, len(callQueryResponse.Result), 2)
	assert.Equal(t, callQueryResponse.Result[0].Caller, "123")
	assert.Equal(t, callQueryResponse.Result[0].IsInbound, false)
	assert.Equal(t, callQueryResponse.Result[1].IsInbound, true)
}
