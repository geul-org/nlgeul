//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestJudge — the verdict comes back; with a trace, a line for a pass and a line with the reason for a rejection; without a trace, nothing is noted
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
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
	sc.trace = nil
	if v := sc.judge(e, 3, 4, rule.Piece{HasStem: true, Prev: '으', Next: -1}); v != 1 {
		t.Fatal(v)
	}
}
