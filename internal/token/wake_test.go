//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestWake — the readings within the gap of the best one are awake, seven at most; the others, and those past the seventh, are asleep; no readings give none
package token

import "testing"

func TestWake(t *testing.T) {
	near := []*path{{verdict: 1, pref: -1}, {verdict: 1, pref: -2}, {verdict: 1, pref: -1 - wakeGap - 1}}
	awake, asleep := wake(near)
	if len(awake) != 2 || len(asleep) != 1 || asleep[0] != near[2] {
		t.Fatal(awake, asleep)
	}
	var many []*path
	for i := 0; i < branches+2; i++ {
		many = append(many, &path{verdict: 1, pref: -1})
	}
	if awake, asleep = wake(many); len(awake) != branches || len(asleep) != 2 {
		t.Fatal(len(awake), len(asleep))
	}
	if awake, asleep = wake(nil); awake != nil || asleep != nil {
		t.Fatal("no readings")
	}
}
