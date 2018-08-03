package accesslib

import "time"
import "fmt"

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

//Access rate controller should be defended by mutex (at least if we want to
//implement lazy initialization
func AccessRateControl(clientId string) bool {
        limit, ok := ReadLimit(clientId)
        returnValue := true
        if !ok {
            returnValue = false
        } else {
            rateMutex.Lock();
            defer rateMutex.Unlock()
            val, ok := rateLimitMap[clientId]	
            if !ok {
		fmt.Printf("We initiate rate limit for %s\n",clientId);
		val = &accessTrackingStruct{1, time.Now()}
		rateLimitMap[clientId] = val
            } else {
		currTime := time.Now()
		dur := currTime.Sub(val.firstIncomeTime)
		if (dur.Nanoseconds()/1000000 > 1000) {
			val.firstIncomeTime = currTime
			val.incomedRequests = 1
		} else {
                        if limit >= val.incomedRequests {
                            val.incomedRequests++
                        }
			//Here should be limit check
			returnValue = !(limit < val.incomedRequests)
		}
            }
            fmt.Printf("Now time is %s, request %d\n", val.firstIncomeTime, val.incomedRequests);
        }
	return returnValue;
}
