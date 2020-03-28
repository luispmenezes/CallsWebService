package model

type CallQueryResult struct {
	Page       int
	TotalPages int
	PageSize   int
	Result     []Call
}
