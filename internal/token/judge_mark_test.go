//ff:func feature=token type=rule control=sequence
//ff:what TestJudgeMark — the candidates that stand come back one by one, each with its own verdict; each is noted in the trace, a defeated one with its reason; without a trace nothing is noted
package token

import (
	"strings"
	"testing"
)

func TestJudgeMark(t *testing.T) {
	sc := scanOf("x")
	candidates := sc.judgeMark(sc.ix.joins, dotFacts([]string{"task", "py"}, 0), ".")
	if len(candidates) != 2 || candidates[0].name != "파일명:." || candidates[1].name != "도메인:." || candidates[0].v != 1 || candidates[1].v != 1 || len(*sc.trace) != 4 || !strings.Contains((*sc.trace)[2], "기각: 양쪽이 맞는 글자 열이 아니다") {
		t.Fatal(candidates, *sc.trace)
	}
	sc.trace = nil
	if candidates = sc.judgeMark(sc.ix.joins, dotFacts([]string{"a", "b"}, 0), "."); len(candidates) != 0 {
		t.Fatal(candidates)
	}
}
