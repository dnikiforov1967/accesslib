package accesslib

import (
	"time"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func sleep(ms time.Duration) {
	timer := time.NewTimer(ms*time.Millisecond)
	<- timer.C
}

func TestSlowRate(t *testing.T) {
	ClientLimits["A"]=3
	res := AccessRateControl("A")
	assertEqual(t, res, true)

	sleep(300)

	res = AccessRateControl("A")
	assertEqual(t, res, true)

	sleep(300)

	res = AccessRateControl("A")
	assertEqual(t, res, true)

	sleep(300)

	res = AccessRateControl("A")
	assertEqual(t, res, false)

	sleep(300)

	res = AccessRateControl("A")
	assertEqual(t, res, true)
}

func readOutside(clientId string, c chan int64) {
    res, _ := ReadLimit(clientId)
    c <- res
}

func TestRLocks(t *testing.T) {
    ClientLimits["A"]=3
    limitMutex.RLock()

    res, _ := ReadLimit("A")
    assertEqual(t, res, int64(3))

    go WriteLimit("A",5)
    timer := time.NewTimer(5*time.Second)
    <-timer.C
    //No repeated call from same thread !
    res = ClientLimits["A"]
    assertEqual(t, res, int64(3))
    ch := make(chan int64)
    go readOutside("A",ch)

    limitMutex.RUnlock()

    limitMutex.RLock()
    res, _ = ReadLimit("A")
    assertEqual(t, res, int64(5))
    limitMutex.RUnlock()
    
    res = <-ch
    assertEqual(t, res, int64(5))
} 
