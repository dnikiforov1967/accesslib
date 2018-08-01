package accesslib

import "time"

type AccessLimitStruct struct {
	ClientId string `json:"clientId"`
	Limit int64 `json:"limit"`
}

type accessTrackingStruct struct {
	incomedRequests int64
	firstIncomeTime time.Time
}

