package accesslib

import (
	"time"
	"testing"
        "github.com/stretchr/testify/assert"
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
	ClientLimits.WriteLimit("A",3)
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

func TestAccessLimit(t *testing.T) {
    ClientLimits.WriteLimit("A",10)
    for i :=0; i <=25; i++ {
	go AccessRateControl("A")
    }
    timer := time.NewTimer(2 * time.Second)
    <-timer.C
    str, _ := rateTracking.readRateMap("A")
    x, _ := str.readTrackingInfo()
    var exp int64 = 11
    assert.Equal(t, exp, x, "Incorrect concurrency control")
    
}