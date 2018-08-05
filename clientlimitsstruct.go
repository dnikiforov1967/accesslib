package accesslib

import "sync"

type clientLimitsStruct struct {
    clientLimits map[string]int64
    limitMutex sync.RWMutex
}

func (obj *clientLimitsStruct) ReadLimit(clientId string) (int64, bool) {
	obj.limitMutex.RLock()
	defer obj.limitMutex.RUnlock()
	limit, ok := obj.clientLimits[clientId]
	return limit, ok
}

func (obj *clientLimitsStruct) ReadLimits() map[string]int64 {
	obj.limitMutex.RLock()
	defer obj.limitMutex.RUnlock()
	results := make(map[string]int64)
	for key, value := range obj.clientLimits {
		results[key] = value
	}
	return results
}

func (obj *clientLimitsStruct) WriteLimit(clientId string, limit int64) {
    obj.limitMutex.Lock()
    defer obj.limitMutex.Unlock()
    obj.clientLimits[clientId] = limit
}