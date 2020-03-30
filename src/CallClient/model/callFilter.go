package model

import (
	"fmt"
	"strconv"
	"time"
)

type Filter struct {
	ParamMap map[string]string
}

func NewFilter() *Filter{
	return &Filter{ParamMap: map[string]string{}}
}

func (f *Filter) SetCaller(caller string){
	f.ParamMap["caller"] = caller
}

func (f *Filter) SetCallee(callee string){
	f.ParamMap["callee"] = callee
}

func (f *Filter) SetStartTime(startTime time.Time){
	f.ParamMap["startTime"] = startTime.Format(time.RFC3339)
}

func (f *Filter) SetEndTime(endTime time.Time){
	f.ParamMap["endTime"] = endTime.Format(time.RFC3339)
}

func (f *Filter) SetInbound(isInbound bool){
	f.ParamMap["inbound"] = strconv.FormatBool(isInbound)
}

func (f *Filter) SetDuration(duration uint16){
	f.ParamMap["duration"] = fmt.Sprintf("%d",duration)
}

func (f *Filter) SetCost(cost uint32){
	f.ParamMap["cost"] = fmt.Sprintf("%d",cost)
}

