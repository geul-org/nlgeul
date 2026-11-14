//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestBoundHead — 것, 것을, 때까지, 거지, 것이다 begin with a dependent noun; 수업을, 사과가, 아는거 and 먹었다 do not
package parse

import "testing"

func TestBoundHead(t *testing.T) {
	for _, s := range []string{"것", "것을", "때까지", "거지", "것이다", "대로"} {
		if !readWord(s, false, &sentence{}).bound {
			t.Errorf("%s", s)
		}
	}
	for _, s := range []string{"수업을", "사과가", "아는거", "먹었다", "매우"} {
		if readWord(s, false, &sentence{}).bound {
			t.Errorf("%s", s)
		}
	}
}
