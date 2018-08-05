package accesslib

import "sync"

var limitMutex sync.RWMutex
var rateMapMutex sync.RWMutex 

var rateLimitMap map[string]*accessTrackingStruct = make(map[string]*accessTrackingStruct)
var ClientLimits map[string]int64 = make(map[string]int64)
