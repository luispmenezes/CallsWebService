package model

type ValidationError struct {
	Id          string
	Description ErrorType
}

type ErrorType string

const (
	INVALID_CALLER_FORMAT ErrorType = "Invalid caller format (digits optionally preceded by country code)"
	INVALID_CALLEE_FORMAT ErrorType = "Invalid callee format (digits optionally preceded by country code)"
	CALLER_EQ_CALLEE      ErrorType = "Invalid caller/callee pair (equal)"
	INVALID_DATE_PAIR     ErrorType = "Invalid date pair (start_date > end_time)"
)
