//ff:func feature=parse type=rule control=sequence
//ff:what TestJudgeSplit — a closing whose graph lifts a closed-set word's exception (split) is judged with the kinds the word table gives the letters before it: 지만 stands in 하지만 after the table's verb 하; with those letters not in the table, and for a closing that is not split, the word's exception holds
package parse

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestJudgeSplit(t *testing.T) {
	sc := scanOf("하지만")
	e := &entry{name: "맺음:지만", form: "지만", class: classClosing}
	markerGraph(e, []string{"하지만"}, nil, true)
	p := rule.Piece{Form: "지만", HasStem: true, Prev: '하', Next: -1}
	if v := sc.judge(e, 1, 3, p); v > 0 {
		t.Fatal("no table gives 하", v)
	}
	sc.ix.lex = map[string]word{"하": {kinds: kindVerb}}
	if v := sc.judge(e, 1, 3, p); v <= 0 {
		t.Fatal("after the table's verb 하", v)
	}
	e.split = false
	if v := sc.judge(e, 1, 3, p); v > 0 {
		t.Fatal("a closing that is not split asks for no kinds", v)
	}
	notes := *sc.trace
	if len(notes) != 3 || !strings.Contains(notes[0], whyWhole) || !strings.HasSuffix(notes[1], "통과") || !strings.Contains(notes[2], whyWhole) {
		t.Fatal(notes)
	}
}
