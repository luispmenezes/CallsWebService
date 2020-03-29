package persistence

import (
	"CallServer/model"
	"time"
)

type Manager interface {
	AddCalls(calls *[]model.Call) error
	RemoveCall(filterParams map[string]interface{}) error
	GetCalls(filterParams map[string]interface{}, pageIdx, pageSize int) (model.CallQueryResult, error)
	GetMetadata(startTime time.Time, endTime time.Time) (model.CallMetadata, error)
}
