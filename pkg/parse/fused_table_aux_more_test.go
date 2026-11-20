//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAuxMore — the linking closings and the range particles added to auxClosings and auxParticles: 부터 after 면서 (가+면서+부터), 조차 after 어서 (먹+어서+조차), 은 after 지만 (가+지만+은), 는 after 려다가 (가+려다가+는) and 만 after 거니 (가+거니+만) are one more piece of the predicate's tail, the closing named as the table names it
package parse

import (
	"slices"
	"testing"
)

func TestAuxMore(t *testing.T) {
	for _, c := range []struct {
		text, closing, particle string
	}{
		{"가면서부터", "맺음:면서", "부터"}, {"먹어서조차", "맺음:어서", "조차"}, {"가지만은", "맺음:지만", "은"},
		{"가려다가는", "맺음:려다가", "는"}, {"가거니만", "맺음:거니", "만"},
	} {
		n := len([]rune(c.text)) - len([]rune(c.particle))
		pc := piece{text: c.particle, names: "보조사:" + c.particle, verdict: 1}
		got := scanOf(c.text).afterClosings(n, 0, pc, &entry{name: "보조사:" + c.particle})
		if !slices.ContainsFunc(got, func(l tail) bool {
			n := len(l.pieces)
			return n == 2 && l.pieces[0].names == c.closing && l.pieces[1].names == "보조사:"+c.particle
		}) {
			t.Errorf("%s: %+v", c.text, got)
		}
	}
}
