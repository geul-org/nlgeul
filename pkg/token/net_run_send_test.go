//ff:func feature=token type=reader control=sequence
//ff:what TestNetRunSend — the places worked out when the first stretch of a net is sent are noted, once, and the stretches are counted
package token

import "testing"

func TestNetRunSend(t *testing.T) {
	r, _ := windowSteps(t, 1, 4)
	r.send(unitStretch{})
	r.step()
	r.send(unitStretch{})
	if r.made != 2 || len(r.stats.firsts) != 1 || r.stats.firsts[0] != 4 {
		t.Fatal(r.made, r.stats.firsts)
	}
}
