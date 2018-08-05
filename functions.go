package accesslib

import (
    "time"
    "sync"
)

func ReadLimit(clientId string) (int64, bool) {
	limitMutex.RLock()
	defer limitMutex.RUnlock()
	limit, ok := ClientLimits[clientId]
	return limit, ok
}

func ReadLimits() map[string]int64 {
	limitMutex.RLock()
	defer limitMutex.RUnlock()
	results := make(map[string]int64)
	for key, value := range ClientLimits {
		results[key] = value
	}
	return results
}

func WriteLimit(clientId string, limit int64) {
	limitMutex.Lock()
	defer limitMutex.Unlock()
	ClientLimits[clientId] = limit
}

func readLimitMap(clientId string) (*accessTrackingStruct, bool){
    rateMapMutex.RLock();
    defer rateMapMutex.RUnlock()
    val, ok := rateLimitMap[clientId]
    return val, ok    
}

func writeLimitMap(clientId string, val *accessTrackingStruct) {
    rateMapMutex.Lock();
    defer rateMapMutex.Unlock()
    rateLimitMap[clientId] = val   
}

//Access rate controller should be defended by mutex (at least if we want to
//implement lazy initialization
func AccessRateControl(clientId string) bool {
        limit, ok := ReadLimit(clientId)
        returnValue := true
        if !ok {
            returnValue = false
        } else {
            val, ok := readLimitMap(clientId)
            if !ok {
		val = &accessTrackingStruct{1, time.Now(), sync.RWMutex{}}
		writeLimitMap(clientId, val)
            } else {
		currTime := time.Now()
                requests, timestamp := val.readTrackingInfo()
		dur := currTime.Sub(timestamp)
		if (dur.Nanoseconds()/1000000 > 1000) {
			val.writeTrackingInfo(1,currTime)
		} else {
                    if limit >= requests {
                        requests = val.incrementIncomeRequests()
                    }
                    //Here should be limit check
                    returnValue = !(limit < requests)
		}
            }
        }
	return returnValue;
}
