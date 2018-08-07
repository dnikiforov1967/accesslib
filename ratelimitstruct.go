package accesslib

import (
    "sync"
    "time"
)

type rateTrackingStruct struct {
    rateTrackingMap map[string]*accessTrackingStruct
    rateMapMutex sync.RWMutex
}

func (obj *rateTrackingStruct) readRateMap(clientId string) (*accessTrackingStruct, bool){
    obj.rateMapMutex.RLock();
    defer obj.rateMapMutex.RUnlock()
    val, ok := obj.rateTrackingMap[clientId]
    return val, ok    
}

func (obj *rateTrackingStruct) initRateMap(clientId string) *accessTrackingStruct {
    obj.rateMapMutex.Lock();
    defer obj.rateMapMutex.Unlock()
    res, ok := obj.rateTrackingMap[clientId]
    if !ok {
        res = &accessTrackingStruct{1, time.Now(), sync.RWMutex{}}
        obj.rateTrackingMap[clientId] = res
    }
    return res
       
}