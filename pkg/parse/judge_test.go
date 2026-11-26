//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestJudge — the verdict comes back; with a trace, a line for a pass and a line with the reason for a rejection; a fused piece is written in brackets; without a trace, nothing is noted; a closing that is split is handed the kinds the word table gives the letters before it
package parse

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestJudge(t *testing.T) {
	sc := scanOf("다음으로")
	var e *entry
	for _, x := range sc.ix.byKey['로'] {
		if x.name == "부사격:로" {
			e = x
		}
	}
	if v := sc.judge(e, 3, 4, rule.Piece{HasStem: true, Prev: '으', Next: -1}); v != 1 {
		t.Fatal(v)
	}
	if v := sc.judge(e, 3, 4, rule.Piece{HasStem: true, Prev: '음', Next: -1}); v != -1 {
		t.Fatal(v)
	}
	notes := *sc.trace
	if len(notes) != 2 || !strings.HasSuffix(notes[0], "통과") || !strings.Contains(notes[1], "기각: 앞 토큰의 끝이 맞지 않는다") {
		t.Fatal(notes)
	}
	fused := &entry{name: "맺음:ㄴ다", form: "다", lead: rule.JongN, g: e.g}
	scanOf("마신다").judge(fused, 2, 3, rule.Piece{HasStem: true, Next: -1})
	sc2 := scanOf("마신다")
	sc2.judge(fused, 2, 3, rule.Piece{HasStem: true, Next: -1})
	if !strings.HasPrefix((*sc2.trace)[0], "  · (ㄴ)다\t") {
		t.Fatal(*sc2.trace)
	}
	sc.trace = nil
	if v := sc.judge(e, 3, 4, rule.Piece{HasStem: true, Prev: '으', Next: -1}); v != 1 {
		t.Fatal(v)
	}
	split := &entry{name: "맺음:지만", form: "지만", class: classClosing}
	markerGraph(split, []string{"하지만"}, nil, true)
	sc3 := scanOf("하지만")
	sc3.ix.lex = map[string]word{"하": {kinds: kindVerb}}
	if v := sc3.judge(split, 1, 3, rule.Piece{Form: "지만", HasStem: true, Prev: '하', Next: -1}); v <= 0 {
		t.Fatal("a split closing is judged with the kinds of the letters before it", v)
	}
}
