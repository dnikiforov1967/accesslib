package accesslib

import "sync"


var ClientLimits clientLimitsStruct = clientLimitsStruct{make(map[string]int64), sync.RWMutex{}}

var rateTracking rateTrackingStruct = rateTrackingStruct{make(map[string]*accessTrackingStruct), sync.RWMutex{}}
