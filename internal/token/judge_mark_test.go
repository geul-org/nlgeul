//ff:func feature=token type=rule control=sequence
//ff:what TestJudgeMark — the readings that stand come back joined with their best verdict; each is noted in the trace, a defeated one with its reason; without a trace nothing is noted
package token

import (
	"strings"
	"testing"
)

func TestJudgeMark(t *testing.T) {
	sc := scanOf("x")
	names, v := sc.judgeMark(sc.ix.joins, dotFacts([]string{"task", "py"}, 0), ".")
	if names != "파일명:.|도메인:." || v != 1 || len(*sc.trace) != 4 || !strings.Contains((*sc.trace)[2], "기각: 양쪽이 맞는 글자 열이 아니다") {
		t.Fatal(names, v, *sc.trace)
	}
	sc.trace = nil
	if names, v = sc.judgeMark(sc.ix.joins, dotFacts([]string{"a", "b"}, 0), "."); names != "" || v > 0 {
		t.Fatal(names, v)
	}
}
