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
