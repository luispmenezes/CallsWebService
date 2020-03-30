package model

import "time"

type Call struct {
	Caller    string
	Callee    string
	StartTime time.Time
	EndTime   time.Time
	IsInbound bool
	Duration  uint16
	Cost      uint32
}

type CallQueryResult struct {
	Page       int
	TotalPages int
	PageSize   int
	Result     []Call
}

