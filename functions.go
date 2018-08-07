package accesslib

import (
    "time"
)

//Access rate controller should be defended by mutex (at least if we want to
//implement lazy initialization
func AccessRateControl(clientId string) bool {
        limit, ok := ClientLimits.ReadLimit(clientId)
        returnValue := true
        if !ok {
            returnValue = false
        } else {
            val, ok := rateTracking.readRateMap(clientId)
            if !ok {
		val = rateTracking.initRateMap(clientId)
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
