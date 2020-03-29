package model

import "time"

type CallMetadata struct {
	tableName             struct{}          `pg:"callws.call_metadata"`
	StartTime             time.Time         `pg:"start_time,pk"`
	EndTime               time.Time         `pg:"end_time,pk"`
	TotalInboundDuration  uint32            `pg:"total_inbound_duration,use_zero"`
	TotalOutboundDuration uint32            `pg:"total_outbound_duration,use_zero"`
	TotalCalls            uint32            `pg:"total_calls,use_zero"`
	TotalCallCost         uint64            `pg:"total_call_costs,use_zero"`
	CallsByCaller         map[string]uint32 `pg:"calls_by_caller"`
	CallsByCallee         map[string]uint32 `pg:"calls_by_callee"`
}

type MetadataQueryResult struct {
	StartTime time.Time `pg:"start_time"`
	Caller    string    `pg:"caller"`
	Callee    string    `pg:"callee"`
	Inbound   bool      `pg:"inbound"`
	Count     uint16    `pg:"count"`
	Duration  uint16    `pg:"total_duration"`
	Cost      uint32    `pg:"total_cost"`
}