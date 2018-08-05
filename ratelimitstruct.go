package accesslib

import "sync"

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

func (obj *rateTrackingStruct) writeRateMap(clientId string, val *accessTrackingStruct) {
    obj.rateMapMutex.Lock();
    defer obj.rateMapMutex.Unlock()
    obj.rateTrackingMap[clientId] = val   
}