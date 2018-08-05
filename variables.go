package accesslib

import "sync"

var limitMutex sync.RWMutex
var ClientLimits map[string]int64 = make(map[string]int64)

var rateTracking rateTrackingStruct = rateTrackingStruct{make(map[string]*accessTrackingStruct), sync.RWMutex{}}
