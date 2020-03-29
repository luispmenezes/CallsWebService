package persistence

import (
	"CallServer/model"
)

type Manager interface {
	AddCalls(calls *[]model.Call) error
	RemoveCall(filterParams map[string]interface{}) error
	GetCalls(filterParams map[string]interface{}, pageIdx, pageSize int) (model.CallQueryResult, error)
	GetMetadata() ([]model.CallMetadata, error)
}
