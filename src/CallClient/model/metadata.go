package model

import "time"

type CallMetadata struct {
	Day                   time.Time
	TotalInboundDuration  uint32
	TotalOutboundDuration uint32
	TotalCalls            uint32
	TotalCallCost         uint64
	CallsByCaller         map[string]uint32
	CallsByCallee         map[string]uint32
}