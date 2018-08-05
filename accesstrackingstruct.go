package accesslib

import (
    "time"
    "sync"
)

type accessTrackingStruct struct {
    incomedRequests int64
    firstIncomeTime time.Time
    innerLock sync.RWMutex    
}

func (obj *accessTrackingStruct) readTrackingInfo() (int64, time.Time) {
    obj.innerLock.RLock()
    defer obj.innerLock.RUnlock()
    return obj.incomedRequests, obj.firstIncomeTime
}

func (obj *accessTrackingStruct) writeTrackingInfo(cnt int64, timestamp time.Time) {
    obj.innerLock.Lock()
    defer obj.innerLock.Unlock()
    obj.incomedRequests = cnt
    obj.firstIncomeTime = timestamp
}

func (obj *accessTrackingStruct) incrementIncomeRequests() int64 {
    obj.innerLock.Lock()
    defer obj.innerLock.Unlock()
    obj.incomedRequests++
    return obj.incomedRequests
}