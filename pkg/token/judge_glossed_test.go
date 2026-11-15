//ff:func feature=token type=rule control=sequence
//ff:what TestJudgeGlossed — judge tells the graph whether the unit is glossed with Han characters: a particle that stands in a plain unit is turned down in a glossed one, with the reason in the trace (a fused piece written in brackets too), and without a trace as well; the piece handed in is not told so by the caller
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestJudgeGlossed(t *testing.T) {
	sc := scanOf("다음으로")
	e := sc.ix.named('로', "부사격:로")
	p := rule.Piece{HasStem: true, Prev: '으', Next: -1}
	if v := sc.judge(e, 3, 4, p); v != 1 {
		t.Fatal("a plain unit", v)
	}
	sc.glossed = true
	if v := sc.judge(e, 3, 4, p); v != -1 {
		t.Fatal("a glossed unit", v)
	}
	notes := *sc.trace
	if len(notes) != 2 || !strings.HasSuffix(notes[0], "통과") || !strings.HasSuffix(notes[1], "기각: "+whyGlossed) {
		t.Fatal(notes)
	}
	fused := &entry{name: "맺음:ㄴ다", form: "다", lead: rule.JongN, g: e.g}
	sf := scanOf("마신다")
	sf.glossed = true
	sf.judge(fused, 2, 3, rule.Piece{HasStem: true, Prev: '으', Next: -1})
	if !strings.HasPrefix((*sf.trace)[0], "  · (ㄴ)다\t") || !strings.Contains((*sf.trace)[0], whyGlossed) {
		t.Fatal(*sf.trace)
	}
	sc.trace = nil
	if v := sc.judge(e, 3, 4, p); v != -1 {
		t.Fatal("a glossed unit with no trace", v)
	}
	sc.glossed = false
	if v := sc.judge(e, 3, 4, p); v != 1 {
		t.Fatal("a plain unit with no trace (the answer remembered for the glossed one is not given)", v)
	}
}
