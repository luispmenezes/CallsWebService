package model

import "time"

type CallMetadata struct {
	tableName     struct{}  `pg:"callws.call_metadata"`
	StartTime     time.Time `pg:"start_time,pk"`
	InboundCalls  int64     `pg:"inboundCalls,use_zero"`
	OutboundCalls int64     `pg:"outboundCalls,use_zero"`
	TotalCalls    int64     `pg:"total_calls,use_zero"`
	TotalCallCost int64     `pg:"total_call_costs,use_zero"`
	CallsByCaller bool      `pg:"calls_by_caller"`
	CallsByCallee bool      `pg:"calls_by_callee"`
}
