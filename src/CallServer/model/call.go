package model

import (
	"CallServer/config"
	"regexp"
	"time"
)

type Call struct {
	tableName struct{}  `pg:"callws.call_data"`
	Caller    string    `pg:"caller,pk"`
	Callee    string    `pg:"callee"`
	StartTime time.Time `pg:"start_time,pk"`
	EndTime   time.Time `pg:"end_time"`
	IsInbound bool      `pg:"inbound,use_zero"`
	Duration  uint16    `pg:"duration,use_zero"`
	Cost      uint32    `pg:"call_cost,use_zero"`
}

func (c *Call) ComputeDurationAndCost() {
	c.Duration = uint16(c.EndTime.Sub(c.StartTime).Minutes())

	callCost := config.GetConfiguration().Server.CallCost

	if c.IsInbound {
		if c.Duration > callCost.InboundPriceThreshold {
			c.Cost = (uint32(callCost.InboundPriceThreshold) * callCost.InboundPrice1) +
				uint32(c.Duration-callCost.InboundPriceThreshold)*callCost.InboundPrice2
		} else {
			c.Cost = uint32(c.Duration) * callCost.InboundPrice1
		}
	} else {
		if c.Duration > callCost.OutboundPriceThreshold {
			c.Cost = (uint32(callCost.OutboundPriceThreshold) * callCost.OutboundPrice1) +
				uint32(c.Duration-callCost.OutboundPriceThreshold)*callCost.OutboundPrice2
		} else {
			c.Cost = uint32(c.Duration) * callCost.OutboundPrice1
		}
	}

}

func (c *Call) Validate() []ValidationError {
	var errorList []ValidationError
	var callerFormatRegex = regexp.MustCompile(config.GetConfiguration().Server.PhoneNumberRegex)

	callId := c.Caller + "-" + c.Callee + "-" + c.StartTime.String()

	if !callerFormatRegex.MatchString(c.Caller) {
		errorList = append(errorList, ValidationError{
			Id:          callId,
			Description: INVALID_CALLER_FORMAT,
		})
	}

	if !callerFormatRegex.MatchString(c.Caller) {
		errorList = append(errorList, ValidationError{
			Id:          callId,
			Description: INVALID_CALLEE_FORMAT,
		})
	}

	if c.Caller == c.Callee {
		errorList = append(errorList, ValidationError{
			Id:          callId,
			Description: CALLER_EQ_CALLEE,
		})
	}

	if !c.EndTime.IsZero() && c.EndTime.Before(c.StartTime) {
		errorList = append(errorList, ValidationError{
			Id:          callId,
			Description: INVALID_DATE_PAIR,
		})
	}

	return errorList
}
