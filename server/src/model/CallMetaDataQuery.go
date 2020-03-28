package model

type CallMetaDataQuery struct {
	Inbound       bool
	Caller        string
	Callee        string
	Count         int
	TotalDuration uint16
	TotalCost     uint32
}
