package api

import (
	"CallServer/model"
	"encoding/json"
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

func (m *MetaDataMockPersistenceManager) RemoveCall(filterParams map[string]interface{}) (int, error) {
	panic("not implemented")
}

func (m *MetaDataMockPersistenceManager) GetCalls(filterParams map[string]interface{}, pageIdx, pageSize int) (model.CallQueryResult, error) {
	panic("not implemented")
}

func (m *MetaDataMockPersistenceManager) GetMetadata() ([]model.CallMetadata, error) {
	return []model.CallMetadata{{
		Day:                   time.Time{},
		TotalInboundDuration:  1,
		TotalOutboundDuration: 2,
		TotalCalls:            3,
		TotalCallCost:         4,
		CallsByCaller:         map[string]uint32{},
		CallsByCallee:         map[string]uint32{},
	}}, nil
}

func performMetaRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	return rec
}

func TestBaseController_GetCallMetadata(t *testing.T) {
	baseController := NewBaseController(&MetaDataMockPersistenceManager{})
	baseController.initializeRoutes()

	response := performMetaRequest(baseController.Engine, "GET", "/metadata")

	assert.Equal(t, http.StatusOK, response.Code)
	var metaDataResponse []model.CallMetadata
	err := json.Unmarshal([]byte(response.Body.String()), &metaDataResponse)

	assert.Nil(t, err)
	assert.Equal(t, len(metaDataResponse), 1)
	assert.Equal(t, metaDataResponse[0].TotalInboundDuration, uint32(1))
	assert.Equal(t, metaDataResponse[0].TotalOutboundDuration, uint32(2))
	assert.Equal(t, metaDataResponse[0].TotalCalls, uint32(3))
	assert.Equal(t, metaDataResponse[0].TotalCallCost, uint64(4))
}
