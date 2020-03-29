package persistence

import (
	"CallServer/model"
	"fmt"
	"github.com/go-pg/pg/v9"
	"log"
	"math"
	"time"
)

const DefaultPageSize = 50

type Manager struct {
	Database *pg.DB
}

func NewManager(host, port, database, username, password string) *Manager {
	log.Printf("Starting Persistence manager (host:%s, port:%s, db:%s user:%s)", host, port, database, username)
	return &Manager{Database: pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Database: database,
		User:     username,
		Password: password,
	})}
}

func (pm *Manager) AddCalls(calls *[]model.Call) error {
	return pm.Database.Insert(calls)
}

func (pm *Manager) RemoveCall(filterParams map[string]interface{}) error {
	query := pm.Database.Model(&model.Call{})

	if len(filterParams) > 0 {
		for key, value := range filterParams {
			query = query.Where(key+" = ?", value)
		}
	} else {
		query = query.Where("TRUE")
	}

	res, err := query.Delete()

	if err == nil && res != nil {
		log.Printf("Removed %d calls", res.RowsAffected())
	}

	return err
}

func (pm *Manager) GetCalls(filterParams map[string]interface{}, pageIdx, pageSize int) (model.CallQueryResult, error) {
	if pageSize == 0 || pageSize > DefaultPageSize {
		pageSize = DefaultPageSize
	}

	recordCount, err := pm.Database.Model((*model.Call)(nil)).Count()
	totalPages := int(math.Ceil(float64(recordCount) / float64(pageSize)))

	if pageIdx >= totalPages {
		return model.CallQueryResult{
			Page:       pageIdx,
			TotalPages: totalPages,
			PageSize:   pageSize,
			Result:     []model.Call{},
		}, nil
	}

	var calls []model.Call
	query := pm.Database.Model(&calls).Order("start_time ASC")

	for key, value := range filterParams {
		query = query.Where(key+" = ?", value)
	}

	err = query.Limit(pageSize).Offset(pageIdx).Select()

	return model.CallQueryResult{
		Page:       pageIdx,
		TotalPages: totalPages,
		PageSize:   pageSize,
		Result:     calls,
	}, err
}

func (pm *Manager) GetMetadata(startTime time.Time, endTime time.Time) (model.CallMetadata, error) {
	callMetadata := model.CallMetadata{}
	var queryResults []model.MetadataQueryResult

	err := pm.Database.Model(&model.Call{}).
		Column("caller", "callee", "inbound").
		ColumnExpr("COUNT(*) as count").
		ColumnExpr("SUM(duration) as total_duration").
		ColumnExpr("SUM(call_cost) as total_cost").
		Where("start_time >= ?", startTime).
		Where("end_time <= ?", endTime).
		Group("caller", "callee", "inbound").
		Select(&queryResults)

	if err != nil {
		return callMetadata, err
	}

	log.Println(queryResults)

	callMetadata.StartTime = startTime
	callMetadata.EndTime = endTime
	callMetadata.CallsByCaller = map[string]uint32{}
	callMetadata.CallsByCallee = map[string]uint32{}

	for _, result := range queryResults {
		if result.Inbound {
			callMetadata.TotalInboundDuration += uint32(result.Duration)
		} else {
			callMetadata.TotalOutboundDuration += uint32(result.Duration)
		}

		callMetadata.TotalCalls += uint32(result.Count)
		callMetadata.TotalCallCost += uint64(result.Cost)

		callMetadata.CallsByCaller[result.Caller] += uint32(result.Count)
		callMetadata.CallsByCallee[result.Callee] += uint32(result.Count)
	}

	return callMetadata, nil
}
