package persistence

import (
	"CallServer/model"
	"fmt"
	"github.com/go-pg/pg/v9"
	"log"
	"math"
)

const DefaultPageSize = 50

type PGManager struct {
	Database *pg.DB
}

func NewPGManager(host, port, database, username, password string) *PGManager {
	log.Printf("Starting Persistence manager (host:%s, port:%s, db:%s user:%s)", host, port, database, username)
	return &PGManager{Database: pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Database: database,
		User:     username,
		Password: password,
	})}
}

func (pm *PGManager) AddCalls(calls *[]model.Call) error {
	return pm.Database.Insert(calls)
}

func (pm *PGManager) RemoveCall(filterParams map[string]interface{}) (int,error) {
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
		return res.RowsAffected(), err
	}

	return 0,err
}

func (pm *PGManager) GetCalls(filterParams map[string]interface{}, pageIdx, pageSize int) (model.CallQueryResult, error) {
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

	calls := []model.Call{}
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

func (pm *PGManager) GetMetadata() ([]model.CallMetadata, error) {
	result := []model.CallMetadata{}
	var queryResults []model.MetadataQueryResult

	err := pm.Database.Model(&model.Call{}).
		Column("start_time", "caller", "callee", "inbound").
		ColumnExpr("COUNT(*) as count").
		ColumnExpr("SUM(duration) as total_duration").
		ColumnExpr("SUM(call_cost) as total_cost").
		Group("start_time", "caller", "callee", "inbound").
		Order("start_time ASC").
		Select(&queryResults)

	if err != nil {
		return result, err
	}

	log.Println(queryResults)

	if len(queryResults) > 0 {
		metaData := model.CallMetadata{
			StartTime:     queryResults[0].StartTime,
			EndTime:       queryResults[0].StartTime,
			CallsByCaller: map[string]uint32{},
			CallsByCallee: map[string]uint32{},
		}

		for _, queryEntry := range queryResults {
			if queryEntry.Inbound {
				metaData.TotalInboundDuration += uint32(queryEntry.Duration)
			} else {
				metaData.TotalOutboundDuration += uint32(queryEntry.Duration)
			}

			metaData.TotalCalls += uint32(queryEntry.Count)
			metaData.TotalCallCost += uint64(queryEntry.Cost)

			metaData.CallsByCaller[queryEntry.Caller] += uint32(queryEntry.Count)
			metaData.CallsByCallee[queryEntry.Callee] += uint32(queryEntry.Count)

			if metaData.StartTime.Sub(queryEntry.StartTime).Hours() > 24 {
				result = append(result, metaData)
				metaData = model.CallMetadata{
					StartTime:     metaData.StartTime,
					EndTime:       queryResults[0].StartTime,
					CallsByCaller: map[string]uint32{},
					CallsByCallee: map[string]uint32{},
				}
			} else {
				metaData.EndTime = queryEntry.StartTime
			}
		}
		result = append(result, metaData)
	}

	return result, nil
}
