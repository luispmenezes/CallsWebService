package api

import (
	"CallServer/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MetaDataMockPersistenceManager struct {
}

func (m *MetaDataMockPersistenceManager) AddCalls(calls *[]model.Call) error {
	panic("not implemented")
}

func (m *MetaDataMockPersistenceManager) RemoveCall(filterParams map[string]interface{}) error {
	panic("not implemented")
}

func (m *MetaDataMockPersistenceManager) GetCalls(filterParams map[string]interface{}, pageIdx, pageSize int) (model.CallQueryResult, error) {
	panic("not implemented")
}

func (m *MetaDataMockPersistenceManager) GetMetadata(startTime time.Time, endTime time.Time) (model.CallMetadata, error) {
	if startTime.IsZero() || endTime.IsZero() || endTime.Before(startTime) {
		return model.CallMetadata{}, errors.New("invalid time interval")
	}
	return model.CallMetadata{
		StartTime:             startTime,
		EndTime:               endTime,
		TotalInboundDuration:  1,
		TotalOutboundDuration: 2,
		TotalCalls:            3,
		TotalCallCost:         4,
		CallsByCaller:         map[string]uint32{},
		CallsByCallee:         map[string]uint32{},
	}, nil
}

func performMetaRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	return rec
}

func TestBaseController_GetCallMetadata_NoParams(t *testing.T) {
	baseController := NewBaseController(&MetaDataMockPersistenceManager{})
	baseController.initializeRoutes()

	response := performMetaRequest(baseController.Engine, "GET", "/metadata")

	assert.Equal(t, http.StatusOK, response.Code)
	var metaDataResponse model.CallMetadata
	err := json.Unmarshal([]byte(response.Body.String()), &metaDataResponse)

	assert.Nil(t, err)
	assert.Equal(t, metaDataResponse.EndTime.Sub(metaDataResponse.StartTime), time.Duration(1)*time.Hour)
}

func TestBaseController_GetCallMetadata_InvalidStartTime(t *testing.T) {
	baseController := NewBaseController(&MetaDataMockPersistenceManager{})
	baseController.initializeRoutes()

	response := performMetaRequest(baseController.Engine, "GET", "/metadata?startTime=aaa123aaaa")

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestBaseController_GetCallMetadata_TimeRange(t *testing.T) {
	baseController := NewBaseController(&MetaDataMockPersistenceManager{})
	baseController.initializeRoutes()

	startTime := time.Date(2020, time.January, 1, 1, 1, 1, 1, time.UTC)
	endTime := time.Date(2020, time.January, 2, 1, 1, 1, 1, time.UTC)

	response := performMetaRequest(baseController.Engine, "GET", fmt.Sprintf("/metadata?startTime=%s&endTime=%s",
		startTime.Format(time.RFC3339Nano), endTime.Format(time.RFC3339Nano)))

	assert.Equal(t, http.StatusOK, response.Code)
	var metaDataResponse model.CallMetadata
	err := json.Unmarshal([]byte(response.Body.String()), &metaDataResponse)

	assert.Nil(t, err)
	assert.Equal(t, metaDataResponse.StartTime, startTime)
	assert.Equal(t, metaDataResponse.EndTime, endTime)
	assert.Equal(t, metaDataResponse.TotalInboundDuration, uint32(1))
	assert.Equal(t, metaDataResponse.TotalOutboundDuration, uint32(2))
	assert.Equal(t, metaDataResponse.TotalCalls, uint32(3))
	assert.Equal(t, metaDataResponse.TotalCallCost, uint64(4))
}
